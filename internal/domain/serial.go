package domain

import "time"

type Serial struct {
	Name        string     `json:"name"`
	ReleaseDate *time.Time `json:"release_date"`
	Score       *uint64    `json:"score"`
	Comment     *string    `json:"comment"`
}

type Season struct {
	Number uint64            `json:"number"`
	Series map[string]string `json:"series"`
}
