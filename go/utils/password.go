package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

const (
	Cbcrypt = iota
	Cargon2 // Default
)

const (
	memory      = 64 * 1024
	paralellism = 4
	iterations  = 3
	keyLength   = 32
)

func GenerateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func HashPassword(alg int, pwd string, salt string) (string, error) {
	var hash string

	if alg == Cbcrypt {
		byte_hash, err := bcrypt.GenerateFromPassword([]byte(pwd+salt), 1)
		if err != nil {
			return "", err
		}

		hash = string(byte_hash)
	} else {
		// cargon2
		byte_hash := argon2.Key(
			[]byte(pwd), []byte(salt), iterations, memory, paralellism, keyLength,
		)
		b64hash := base64.RawStdEncoding.EncodeToString(byte_hash)
		hash = fmt.Sprintf(
			"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version,
			memory, iterations, paralellism, salt, b64hash,
		)
	}

	return hash, nil
}

func CheckHash(alg int, pwd string, salt string, hash string) bool {
	var isTrue bool

	if alg == Cbcrypt {
		isTrue = true
	} else {
		isTrue = false
	}
	return isTrue
}
