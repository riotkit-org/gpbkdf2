package main

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodePasswordToPBKDF2_FailsOnUnsupportedAlgorithm(t *testing.T) {
	_, err := EncodePasswordToPBKDF2("password", "salt", 16, 256, "sha2000")
	assert.Contains(t, err.Error(), "Unrecognized hashing algorithm")
}

func TestEncodePasswordToPBKDF2_PrintsValidChoices(t *testing.T) {
	_, err := EncodePasswordToPBKDF2("password", "salt", 16, 256, "sha2000")

	assert.Contains(t, err.Error(), "sha512")
	assert.Contains(t, err.Error(), "sha256")
	assert.Contains(t, err.Error(), "sha1")
	assert.Contains(t, err.Error(), "md5")
}

func TestEncodePasswordToPBKDF2_ProducesValidLengthKey(t *testing.T) {
	keyAsHex, _ := EncodePasswordToPBKDF2("password", "salt", 16, 256, "sha256")
	keyAsBin, _ := hex.DecodeString(keyAsHex)

	assert.Len(t, keyAsBin, 256)
	assert.Len(t, keyAsHex, 512)
}
