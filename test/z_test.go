package test

import (
	"encoding/hex"
	"fmt"
	"simple-crud-app/internal/lib/hash"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestBcrypt(t *testing.T) {
	password := []byte("MyDarkSecret")

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(hashedPassword))
}

func TestSHA256(t *testing.T) {
	s := "ComplexPassword"
	sh, _ := hash.NewSHA256Hash(s)
	t.Log(sh)
}
