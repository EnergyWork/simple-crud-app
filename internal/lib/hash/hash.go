package hash

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"

	errs "simple-crud-app/internal/lib/errors"
)

func NewBcryptHash(data string) (string, *errs.Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
	}
	return hex.EncodeToString(hashedPassword), nil
}

func NewSHA256Hash(data string) (string, *errs.Error) {
	h := sha256.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return "", errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// NewAccessKey returns new uniaque key in hex encoded
func NewAccessKey(password string) (string, *errs.Error) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
	}
	key := argon2.Key([]byte(password), salt, 3, 32*1024, 4, 16)
	return hex.EncodeToString(key), nil
}

func NewSecretKey(password string) (string, *errs.Error) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", errs.New().SetCode(errs.ERROR_INTERNAL).SetMsg("%s", err)
	}
	key := argon2.Key([]byte(password), salt, 3, 32*1024, 4, 32)
	return hex.EncodeToString(key), nil
}
