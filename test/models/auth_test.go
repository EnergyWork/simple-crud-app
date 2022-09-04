package models

import (
	"simple-crud-app/internal/lib/crypto"
	"simple-crud-app/internal/lib/hash"
	"simple-crud-app/internal/models"
	"testing"
	"time"
)

func TestUserRegister(t *testing.T) {
	db := GetConn(t, "../config.yml") // get db connection wuth parameters from config fil
	hashedPassword, _ := hash.NewSHA256Hash("!Qq123456")
	privateKey := crypto.PrivateKey{}
	privateKey.LaodRsaPrivateKey(PRIVATE_KEY)
	a := models.UserAuth{
		UserName:     "Shrek",
		UserPassword: hashedPassword,
		PrivateKey:   privateKey,
		ExpiredAt:    time.Now().AddDate(0, 0, 7),
	}
	if err := a.UserRegister(db); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", a)
}
