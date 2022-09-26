package models

import (
	"simple-crud-app/internal/lib/hash"
	"simple-crud-app/internal/models"
	"testing"
)

func TestUserRegister(t *testing.T) {
	db := GetConn(t, "../config.yml") // get db connection wuth parameters from config fil
	hashedPassword, _ := hash.NewSHA256Hash("!Qq123456")
	a := models.User{
		Login:    "Shrek",
		Password: hashedPassword,
	}
	if err := a.Create(db); err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", a)
}
