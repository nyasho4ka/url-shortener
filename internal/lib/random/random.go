package random

import (
	"crypto/rand"
	"encoding/base64"
)

func NewRandomString(stringLength int) string {
	byteLength := (stringLength*3)/4 + 1
	b := make([]byte, byteLength)

	rand.Read(b)

	encoded := base64.URLEncoding.EncodeToString(b)
	return encoded[:stringLength]
}
