package keystore

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"log"
)

type KeyStore struct {
	PrivateKey string `json:"private_key"`
}

func New(sk string) *KeyStore {
	pvk, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		log.Fatalf("failed to generate private key: %s", err)
	}

	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pvk),
	}

	encryptedBlock, err := x509.EncryptPEMBlock(rand.Reader, pemBlock.Type, pemBlock.Bytes, []byte(sk), x509.PEMCipherAES256)
	blockBytes := pem.EncodeToMemory(encryptedBlock)
	pemStr := hex.EncodeToString(blockBytes)

	if err != nil {
		log.Fatalf("failed to encrypt private key: %s", err)
	}

	k := KeyStore{PrivateKey: pemStr}

	return &k
}

func PrivateKeyFromString(pvkStr, sk string) *rsa.PrivateKey {
	pvkBytes, err := hex.DecodeString(pvkStr)

	if err != nil {
		log.Fatalf("failed to decode private key: %s", err)
	}

	pemBlock, _ := pem.Decode(pvkBytes)
	blockBytes, err := x509.DecryptPEMBlock(pemBlock, []byte(sk))

	if err != nil {
		log.Fatalf("failed to decode private key: %s", err)
	}

	pvk, err := x509.ParsePKCS1PrivateKey(blockBytes)

	if err != nil {
		log.Fatalf("failed to parse private key: %s", err)
	}

	return pvk
}

func PublicKeyFromString(pbkStr string) *rsa.PublicKey {
	pbkBytes, err := hex.DecodeString(pbkStr)

	if err != nil {
		log.Fatalf("failed to decode public sk: %s", err)
	}

	pemBlock, _ := pem.Decode(pbkBytes)

	pbk, err := x509.ParsePKCS1PublicKey(pemBlock.Bytes)

	if err != nil {
		log.Fatalf("failed to parse public sk: %s", err)
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

func EncryptKey(sk string, pbk *rsa.PublicKey) string {
	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pbk, []byte(sk), nil)

	if err != nil {
		log.Fatalf("failed to encrypt private key: %s", err)
	}

	encryptedStr := hex.EncodeToString(encrypted)

	return encryptedStr
}

func DecryptKey(sk string, pvk *rsa.PrivateKey) string {
	encrypted, err := hex.DecodeString(sk)

	if err != nil {
		log.Fatalf("failed to decode private key: %s", err)
	}

	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, pvk, encrypted, nil)

	if err != nil {
		log.Fatalf("failed to decrypt private key: %s", err)
	}

	decryptedStr := string(decrypted)

	return decryptedStr
}
