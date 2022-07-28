package vote

import (
	"crypto/sha512"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/8BITS-COLAB/ballot-box/candidate"
	"github.com/8BITS-COLAB/ballot-box/db"
	"github.com/8BITS-COLAB/ballot-box/keystore"
	"github.com/8BITS-COLAB/ballot-box/voter"
)

type Vote struct {
	Index         int64               `json:"index" gorm:"primaryKey"`
	Timestamp     int64               `json:"timestamp"`
	PrevHash      string              `json:"prev_hash"`
	Hash          string              `json:"hash"`
	VoterAddress  string              `json:"voter_address"`
	Voter         voter.Voter         `json:"voter" gorm:"foreignKey:voter_address"`
	CandidateCode string              `json:"candidate_code"`
	Candidate     candidate.Candidate `json:"candidate" gorm:"foreignKey:candidate_code"`
	Year          int                 `json:"year"`
}

func calculateHash(v Vote) string {
	var nonce int64
	h := sha512.New()
	difficulty := 3

	for {
		s := fmt.Sprintf("%d:%s:%s:%s:%d:%d", v.Index, v.PrevHash, v.VoterAddress, v.CandidateCode, v.Year, nonce)
		h.Write([]byte(s))

		hash := fmt.Sprintf("%x", h.Sum(nil))

		if strings.HasPrefix(hash, strings.Repeat("0", difficulty)) {
			return hash
		}

		nonce++
	}
}

func New(candidateCode string) *Vote {
	var vtr voter.Voter
	var cnd candidate.Candidate

	d, sql := db.New()
	d.AutoMigrate(&Vote{})

	defer sql.Close()

	k := keystore.Show()

	pvkStr := k.PrivateKey

	pvk := keystore.PrivateKeyFromString(pvkStr)

	pbk := pvk.PublicKey
	pemStr := keystore.PublicKeyToString(&pbk)

	if err := d.Where("public_key = ?", pemStr).First(&vtr).Error; err != nil {
		log.Fatalf("failed to get voter: %s", err)
	}

	if err := d.Where("code = ?", candidateCode).First(&cnd).Error; err != nil {
		log.Fatalf("failed to get candidate: %s", err)
	}

	var lv, v Vote

	d.Where("voter_address = ? AND year = ?", vtr.Address, time.Now().Year()).First(&v)

	if v.VoterAddress != "" {
		log.Fatalf("voter already voted")
	}

	d.Order("index desc").First(&lv)

	v.PrevHash = lv.Hash
	v.Index = lv.Index + 1
	v.Timestamp = time.Now().Unix()
	v.VoterAddress = vtr.Address
	v.CandidateCode = cnd.Code
	v.Year = time.Now().Year()

	hash := calculateHash(v)

	v.Hash = hash

	// TODO: SOLVE THIS
	// v.Data = v

	if err := d.Create(&v).Error; err != nil {
		log.Fatalf("failed to create vote: %s", err)
	}

	return &v
}

func Status() map[string]int {
	var vts []Vote
	result := make(map[string]int)

	d, sql := db.New()
	d.AutoMigrate(&Vote{})

	defer sql.Close()

	d.Preload("Candidate").Find(&vts).Where("year = ?", time.Now().Year())

	for _, vt := range vts {
		result[vt.Candidate.Name] += 1
	}

	sort.Slice(vts, func(i, j int) bool {
		return result[vts[i].Candidate.Name] > result[vts[j].Candidate.Name]
	})

	return result
}

func Integrity() interface{} {
	var vts []Vote

	d, sql := db.New()

	defer sql.Close()

	d.Find(&vts)

	for i := 1; i < len(vts); i++ {
		if vts[i].PrevHash != vts[i-1].Hash {
			return fmt.Sprintf("the prev hash of vote %d is not equal to the hash of vote %d. expected: %s, actual: %s", i, i-1, vts[i-1].Hash, vts[i].PrevHash)
		}

		hash := calculateHash(vts[i])

		if hash != vts[i].Hash {
			return fmt.Sprintf("the hash of vote %d is not equal to the hash of vote %d. expected: %s, actual: %s", i, i, hash, vts[i].Hash)
		}
	}

	return "the chain has integrity"
}
