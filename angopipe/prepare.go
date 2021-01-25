package angopipe

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"os"
)

// Prepare 共通処理
func Prepare() (cipher.AEAD, error) {
	rawKey := os.Getenv("ANGO_KEY")
	if rawKey == "" {
		return nil, errors.New("Environment Variable 'ANGO_KEY' is empty")
	}

	key, err := base64.StdEncoding.DecodeString(rawKey)
	if err != nil {
		return nil, errors.New("Decode 'ANGO_KEY' key error")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}
