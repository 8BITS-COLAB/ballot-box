package voter

import (
	"fmt"
	"log"

	"github.com/8BITS-COLAB/ballot-box/db"
	"github.com/8BITS-COLAB/ballot-box/keystore"
)

type Voter struct {
	Address   string `json:"address" gorm:"primaryKey"`
	PublicKey string `json:"public_key"`
	Registry  string `json:"registry" gorm:"uniqueIndex"`
}

func New(registry string, key string) *Voter {
	k := keystore.New(key)

	pvk := keystore.PrivateKeyFromString(k.PrivateKey, key)
	pemStr := keystore.PublicKeyToString(&pvk.PublicKey)

	v := Voter{
		Address:   fmt.Sprintf("%x", pvk.PublicKey.N.Bytes()[:20]),
		PublicKey: pemStr,
		Registry:  registry,
	}

	d, sql := db.New()
	d.AutoMigrate(&Voter{})

	defer sql.Close()

	if err := d.Create(&v).Error; err != nil {
		log.Fatalf("failed to create voter: %s", err)
	}

	return &v
}

func Show(key string) *Voter {
	k := keystore.Show()

	pvk := keystore.PrivateKeyFromString(k.PrivateKey, key)
	pbk := pvk.PublicKey
	pemStr := keystore.PublicKeyToString(&pbk)

	var v Voter

	d, sql := db.New()
	defer sql.Close()

	if err := d.Where("public_key = ?", pemStr).First(&v).Error; err != nil {
		log.Fatalf("failed to get voter: %s", err)
	}

	return &v

}
