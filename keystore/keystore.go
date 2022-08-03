package keystore

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
)

type KeyStore struct {
	PrivateKey string `json:"private_key"`
}

func New(sk string) (*KeyStore, error) {
	pvk, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return nil, err
	}

	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pvk),
	}

	encryptedBlock, err := x509.EncryptPEMBlock(rand.Reader, pemBlock.Type, pemBlock.Bytes, []byte(sk), x509.PEMCipherAES256)

	if err != nil {
		return nil, err
	}

	blockBytes := pem.EncodeToMemory(encryptedBlock)
	pemStr := hex.EncodeToString(blockBytes)

	k := KeyStore{PrivateKey: pemStr}

	return &k, nil
}

func PrivateKeyFromString(pvkStr, sk string) (*rsa.PrivateKey, error) {
	pvkBytes, err := hex.DecodeString(pvkStr)

	if err != nil {
		return nil, err
	}

	pemBlock, _ := pem.Decode(pvkBytes)
	blockBytes, err := x509.DecryptPEMBlock(pemBlock, []byte(sk))

	if err != nil {
		return nil, err
	}

	pvk, err := x509.ParsePKCS1PrivateKey(blockBytes)

	if err != nil {
		return nil, err
	}

	return pvk, nil
}

func PublicKeyFromString(pbkStr string) (*rsa.PublicKey, error) {
	pbkBytes, err := hex.DecodeString(pbkStr)

	if err != nil {
		return nil, err
	}

	pemBlock, _ := pem.Decode(pbkBytes)

	pbk, err := x509.ParsePKCS1PublicKey(pemBlock.Bytes)

	if err != nil {
		return nil, err
	}

	return pbk, nil
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

func EncryptKey(sk string, pbk *rsa.PublicKey) (string, error) {
	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pbk, []byte(sk), nil)

	if err != nil {
		return "", err
	}

	encryptedStr := hex.EncodeToString(encrypted)

	return encryptedStr, nil
}

func DecryptKey(sk string, pvk *rsa.PrivateKey) (string, error) {
	encrypted, err := hex.DecodeString(sk)

	if err != nil {
		return "", err
	}

	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, pvk, encrypted, nil)

	if err != nil {
		return "", err
	}

	decryptedStr := string(decrypted)

	return decryptedStr, nil
}
