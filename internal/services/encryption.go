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
	keyBytes := []byte(key)                 // Convert the key to a byte slice
	decryptedTextBytes := []byte(plainText) // Convert the plain text to a byte slice
	block, err := aes.NewCipher(keyBytes)   // Create a new AES cipher block
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12)                                  // AES-GCM standard nonce size
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil { // Generate a random nonce
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block) // Create a GCM cipher mode instance
	if err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nil, nonce, decryptedTextBytes, nil)              // Encrypt the plain text
	return base64.StdEncoding.EncodeToString(append(nonce, cipherText...)), nil // Return the encrypted text as a base64 string
}

// Decrypt decrypts an AES-GCM encrypted text
func Decrypt(encryptedText string, key string) (string, error) {
	keyBytes := []byte(key)                                     // Convert the key to a byte slice
	data, err := base64.StdEncoding.DecodeString(encryptedText) // Decode the base64 encoded text
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes) // Create a new AES cipher block
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block) // Create a GCM cipher mode instance
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()                            // Get the nonce size
	nonce, cipherText := data[:nonceSize], data[nonceSize:]    // Extract the nonce and cipher text
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil) // Decrypt the cipher text
	if err != nil {
		return "", err
	}

	return string(plainText), nil // Return the decrypted plain text
}
