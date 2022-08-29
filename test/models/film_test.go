package models

import (
	"fmt"
	"math/rand"
	"simple-crud-app/internal/models"
	"testing"
	"time"
)

func TestCreateFilm(t *testing.T) {
	db := GetConn(t, "../config.yml") // get db connection wuth parameters from config fil

	// create model

	// clowning process
	typ := "test type"
	name := "test name"
	releaseDate := time.Date(2000, time.August, 30, 0, 0, 0, 0, time.Local)
	duration := "1h30m"
	serialCount := uint64(rand.Intn(25))
	score := uint64(rand.Intn(100))
	comment := "good movie"
	_ = fmt.Sprint(typ, name, releaseDate, duration, serialCount, score, comment)

	// forming model data
	film := &models.Film{}
	film.Type = &typ
	film.Name = name
	film.ReleaseDate = &releaseDate
	film.Duration = &duration
	film.SerialCount = &serialCount
	//film.Score = score
	//film.Comment = &comment

	// try to create record
	if err := film.Create(db); err != nil {
		t.Fatal(err)
	}
	t.Logf("NEW FILM: %+v", film)
}
