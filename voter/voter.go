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

var d = db.New()

func New(registry string, sk string) (*Voter, *keystore.KeyStore, error) {
	k, err := keystore.New(sk)

	if err != nil {
		return nil, nil, err
	}

	pvk, err := keystore.PrivateKeyFromString(k.PrivateKey, sk)

	if err != nil {
		return nil, nil, err
	}

	pemStr := keystore.PublicKeyToString(&pvk.PublicKey)

	v := Voter{
		Address:   fmt.Sprintf("%x", pvk.PublicKey.N.Bytes()[:20]),
		PublicKey: pemStr,
		Registry:  registry,
	}

	if err := d.Create(&v).Error; err != nil {
		log.Fatalf("failed to create voter: %s", err)
	}

	return &v, k, nil
}

func Show(pvkStr, sk string) (*Voter, error) {
	pvk, err := keystore.PrivateKeyFromString(pvkStr, sk)

	if err != nil {
		return nil, err
	}

	pbk := pvk.PublicKey
	pemStr := keystore.PublicKeyToString(&pbk)

	var v Voter

	if err := d.Where("public_key = ?", pemStr).First(&v).Error; err != nil {
		return nil, err
	}

	return &v, nil

}
