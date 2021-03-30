package foods

import "time"

type Food struct {
	ID          int
	Name        string
	Description string
	Ingredients string
	Slug        string
	// ImageURL    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
