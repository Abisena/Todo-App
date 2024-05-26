package utils

import (
	"encoding/base64"
)

// mengenkripsi string menggunakan algoritma base64
func EncodeString(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// mengdekripsi string yang dienkripsi menggunakan algoritma base64
func DecodeString(s string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}