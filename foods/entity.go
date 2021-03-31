package foods

import "time"

type Food struct {
	ID          int
	Name        string
	Description string
	Ingredients string
	Slug        string
	FoodImages  []FoodImage
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type FoodImage struct {
	ID        int
	FoodID    int
	IsPrimary bool
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
