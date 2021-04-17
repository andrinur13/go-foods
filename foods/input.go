package foods

type FoodInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Ingredients string `json:"ingredients" binding:"required"`
}

type FoodDeelete struct {
	FoodID int `json:"id" binding:"required"`
}

type FoodDetail struct {
	FoodID int `uri:"id" binding:"required"`
}

type FoodEdit struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ingredients string `json:"ingredients"`
}

type FoodImageInput struct {
	FoodID    int  `form:"food_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
}
