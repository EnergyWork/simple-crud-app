package domain

type Film struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	ReleaseDate *int64  `json:"release_date,omitempty"` // Unix
	Duration    *string `json:"duration,omitempty"`
	Score       *uint64 `json:"score,omitempty"`
	Comment     *string `json:"comment,omitempty"`
	Watched     bool    `json:"watched"`
}
