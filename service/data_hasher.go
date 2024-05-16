package service

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
)

type DataHasher struct {
	securityIterations int
	securityKeyLen     int
	saltSize           int
}

func NewDataHasher(saltSize int, securityIterations int, securityKeyLen int) DataHasher {
	return DataHasher{
		saltSize:           saltSize,
		securityIterations: securityIterations,
		securityKeyLen:     securityKeyLen,
	}
}

func (dh DataHasher) HashPassport(password string) string {
	hash := make([]byte, dh.securityKeyLen+dh.saltSize)
	salt := make([]byte, dh.saltSize)
	_, _ = rand.Read(salt)

	hashedPassword := pbkdf2.Key([]byte(password), salt, dh.securityIterations, dh.securityKeyLen, sha256.New)

	copy(hash[0:], salt)
	copy(hash[dh.saltSize:], hashedPassword)

	return base64.StdEncoding.EncodeToString(hash)
}

func (dh DataHasher) VerifyPassport(password, hashedPassword string) bool {
	buf, err := base64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false
	}

	if len(buf) != dh.securityKeyLen+dh.saltSize {
		return false
	}

	salt := make([]byte, dh.saltSize)
	copy(salt, buf[:dh.saltSize])

	hashedPasswordBytes := pbkdf2.Key([]byte(password), salt, dh.securityIterations, dh.securityKeyLen, sha256.New)
	if bytes.Compare(buf[dh.saltSize:], hashedPasswordBytes) != 0 {
		return false
	}

	return true
}
