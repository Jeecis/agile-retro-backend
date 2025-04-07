package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encrypt encrypts plain text using AES-GCM
func Encrypt(plainText string, key string) (string, error) {
	keyBytes := []byte(key)
	decryptedTextBytes := []byte(plainText)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12) // AES-GCM standard nonce size
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nil, nonce, decryptedTextBytes, nil)
	return base64.StdEncoding.EncodeToString(append(nonce, cipherText...)), nil
}

func Decrypt(encryptedText string, key string) (string, error) {
	keyBytes := []byte(key)
	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
