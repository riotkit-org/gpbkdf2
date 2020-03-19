package main

import (
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

func EncodePasswordToPBKDF2(password string, salt string, iterations int, keyLength int, hashingAlgorithm string) string {

	binaryEncoded := pbkdf2.Key([]byte(password), []byte(salt), iterations, keyLength, sha512.New)

	return hex.EncodeToString(binaryEncoded)
}
