package models

import "time"

type Season struct {
	ID        uint64
	SerialID  uint64
	Number    uint64
	Series    map[string]string // {"1": "42m", "2": "46m"}
	CreatedAt time.Time
	UpdatedAt *time.Time
}
