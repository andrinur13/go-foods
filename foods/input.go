package foods

type FoodInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Ingredients string `json:"ingredients" binding:"required"`
}

type FoodImageInput struct {
	FoodID    int  `form:"food_id" binding:"required"`
	IsPrimary bool `form:"is_primary"`
}

type FoodDeleteInput struct {
	FoodID int `json:"id" binding:"required"`
}
