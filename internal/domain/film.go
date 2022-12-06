package domain

type Film struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	ReleaseDate *string `json:"release_date,omitempty"`
	Duration    *string `json:"duration,omitempty"`
	Score       *uint64 `json:"score,omitempty"`
	Comment     *string `json:"comment,omitempty"`
	Watched     bool    `json:"watched"`
}
