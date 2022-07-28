package keystore

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"log"
	"os"

	"github.com/bytedance/sonic"
)

type KeyStore struct {
	PrivateKey string `json:"private_key"`
}

func New() *KeyStore {
	pvk, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}

	pemStr := PrivateKeyToString(pvk)

	k := KeyStore{PrivateKey: pemStr}

	b, err := sonic.Marshal(k)

	if err != nil {
		log.Fatalf("failed to marshal keystore: %s", err)
	}

	f, err := os.Create("keystore.json")

	if err != nil {
		log.Fatalf("failed to create keystore: %s", err)
	}

	_, err = f.Write(b)

	if err != nil {
		log.Fatalf("failed to write keystore: %s", err)
	}

	return &k
}

func Show() *KeyStore {
	var k KeyStore
	f, err := os.Open("keystore.json")

	if err != nil {
		log.Fatalf("failed to open keystore: %s", err)
	}

	if err = sonic.ConfigDefault.NewDecoder(f).Decode(&k); err != nil {
		log.Fatalf("failed to unmarshal keystore: %s", err)
	}

	return &k
}

func PrivateKeyFromString(pvkStr string) *rsa.PrivateKey {
	pvkBytes, err := hex.DecodeString(pvkStr)

	if err != nil {
		log.Fatalf("failed to decode private key: %s", err)
	}

	pemBlock, _ := pem.Decode(pvkBytes)

	if err != nil {
		log.Fatalf("failed to decode private key: %s", err)
	}

	pvk, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)

	if err != nil {
		log.Fatalf("failed to parse private key: %s", err)
	}

	return pvk
}

func PublicKeyFromString(pbkStr string) *rsa.PublicKey {
	pbkBytes, err := hex.DecodeString(pbkStr)

	if err != nil {
		log.Fatalf("failed to decode public key: %s", err)
	}

	pemBlock, _ := pem.Decode(pbkBytes)

	pbk, err := x509.ParsePKCS1PublicKey(pemBlock.Bytes)

	if err != nil {
		log.Fatalf("failed to parse public key: %s", err)
	}

	return pbk
}

func PrivateKeyToString(pvk *rsa.PrivateKey) string {
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pvk),
	}
	pemBytes := pem.EncodeToMemory(pemBlock)
	pemStr := hex.EncodeToString(pemBytes)

	return pemStr
}

func PublicKeyToString(pbk *rsa.PublicKey) string {
	pemBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pbk),
	}
	pemBytes := pem.EncodeToMemory(pemBlock)
	pemStr := hex.EncodeToString(pemBytes)

	return pemStr
}
