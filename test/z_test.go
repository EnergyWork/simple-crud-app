package test

import (
	"encoding/hex"
	"fmt"
	"simple-crud-app/internal/lib/hash"
	"testing"
	"time"

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

func TestTimeAfter(t *testing.T) {
	now := time.Now()
	someTime := time.Date(2022, time.September, now.Day(), 2, 5, 0, 0, time.Local)
	expires := now.After(someTime)
	t.Log(expires)
}

func TestTimeUnix(t *testing.T) {
	var t1 int64 = 1662333342
	var t2 int64 = 1662343993

	t1parse := time.Unix(t1, 0)
	t2parse := time.Unix(t2, 0)

	t.Logf("time1: %s", t1parse)
	t.Logf("time2: %s", t2parse)
	t.Log(t2parse.After(t1parse))
}
