package aesutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

type AESUtil struct {
	key []byte
}

// Encrypt encrypts plaintext using AES encryption.
func (aesUtil *AESUtil) SetKey(key []byte) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		panic("Invalid key size. Only 128-bit, 192-bit, and 256-bit keys are supported.")
	}
	aesUtil.key = key
}

func (aesUtil *AESUtil) validateKey(key []byte) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		panic("Invalid key size. Only 128-bit, 192-bit, and 256-bit keys are supported.")
	}
}

func (aesUtil *AESUtil) Encrypt(plaintext string) (string, error) {
	aesUtil.validateKey(aesUtil.key)
	block, err := aes.NewCipher(aesUtil.key)
	if err != nil {
		return "", fmt.Errorf("Invalid encryption key length: %w", err)
	}

	// Generate a random initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt the plaintext
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, []byte(plaintext))

	// Prepend the IV for use during decryption
	encrypted := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt decrypts ciphertext using AES encryption.
func (aesUtil *AESUtil) Decrypt(encrypted string) (string, error) {
	aesUtil.validateKey(aesUtil.key)
	data, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(aesUtil.key))
	if err != nil {
		return "", err
	}

	// Extract the IV and the ciphertext
	if len(data) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	// Decrypt the ciphertext
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

func NewAESUtil() *AESUtil {
	return &AESUtil{}
}
