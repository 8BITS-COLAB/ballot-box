package candidate

import (
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/8BITS-COLAB/ballot-box/db"
)

type Candidate struct {
	Name  string `json:"name"`
	Code  string `json:"code" gorm:"uniqueIndex;primaryKey"`
	Party string `json:"party"`
}

var d = db.New()

func New(name string, party string) *Candidate {
	h := sha1.New()

	v := fmt.Sprintf("%d%s%s", time.Now().Unix(), name, party)

	h.Write([]byte(v))

	c := Candidate{
		Name:  name,
		Party: party,
		Code:  fmt.Sprintf("%x", h.Sum(nil))[:6],
	}

	if err := d.Create(&c).Error; err != nil {
		log.Fatalf("failed to create candidate: %s", err)
	}

	return &c
}

func All() []Candidate {
	var cs []Candidate

	if err := d.Find(&cs).Error; err != nil {
		log.Fatalf("failed to get candidates: %s", err)
	}

	return cs
}

func GetByCode(code string) (*Candidate, error) {
	var c Candidate

	if err := d.Where("code = ?", code).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}
