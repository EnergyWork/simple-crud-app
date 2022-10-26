package domain

import "time"

type Film struct {
	Name        string     `json:"name"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Duration    *string    `json:"duration,omitempty"`
	Score       *uint64    `json:"score,omitempty"`
	Comment     *string    `json:"comment,omitempty"`
	Watched     bool       `json:"watched"`
}
