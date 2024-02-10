package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// It has to be 32 byte
const secretKey string = "GodIsGoodAllTheTime#31/01/2024#*"

// It has to be 12 byte
const secretNonce string = "apiExempleGo"

func Encrypt(s string) string {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := []byte(secretNonce)
	plaintext := []byte(s)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return base64.URLEncoding.EncodeToString(ciphertext)
}

func Decrypt(s string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(s)
	block, err := aes.NewCipher([]byte(secretKey))

	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := []byte(secretNonce)
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext[:])
}
