package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func getSupportedHashAlgorithms() map[string]func() hash.Hash {
	return map[string]func() hash.Hash{
		"sha512": sha512.New,
		"sha256": sha256.New,
		"sha1":   sha1.New,
		"md5":    md5.New,
	}
}

func getAlgorithmChoices() string {
	elements := getSupportedHashAlgorithms()

	i, choices := 0, ""
	for key, _ := range elements {
		choices += ", " + key
		i++
	}

	return strings.TrimLeft(choices, ", ")
}

func EncodePasswordToPBKDF2(password string, salt string, iterations int, keyLength int, hashingAlgorithm string) (string, error) {
	algo, supportsHashAlgorithm := getSupportedHashAlgorithms()[hashingAlgorithm]

	if !supportsHashAlgorithm {
		return "", errors.New(fmt.Sprintf("Unrecognized hashing algorithm \"%s\". Supported: %s", hashingAlgorithm, getAlgorithmChoices()))
	}

	binaryEncoded := pbkdf2.Key([]byte(password), []byte(salt), iterations, keyLength, algo)

	return hex.EncodeToString(binaryEncoded), nil
}
