package foods

type FoodInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Ingredients string `json:"ingredients" binding:"required"`
}
