package botid

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func Encrypt(password string, data any) (string, error) {
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	iv := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("failed to generate IV: %w", err)
	}

	key, _ := pbkdf2.Key(sha256.New, password, salt, 100000, 32)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data to JSON: %w", err)
	}

	ciphertext := gcm.Seal(nil, iv, jsonData, nil)

	result := make([]byte, 0, len(salt)+len(iv)+len(ciphertext))
	result = append(result, salt...)
	result = append(result, iv...)
	result = append(result, ciphertext...)

	return base64.StdEncoding.EncodeToString(result), nil
}

func Decrypt(password string, encryptedData string) error {
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return fmt.Errorf("failed to decode base64: %w", err)
	}

	if len(data) < 44 {
		return errors.New("encrypted data too short")
	}

	salt := data[:16]

	iv := data[16:28]

	ciphertext := data[28:]

	key, _ := pbkdf2.Key(sha256.New, password, salt, 100000, 32)

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("failed to decrypt data: %w", err)
	}
	fmt.Println(string(plaintext))

	return nil
}
