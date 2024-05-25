package utilities

import (
	"crypto/sha256"
	"encoding/hex"
)

func ToSHA256(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}
