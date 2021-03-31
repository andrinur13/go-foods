package foods

import "strings"

type FoodFormatter struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Ingredients []string `json:"ingredients"`
	Slug        string   `json:"slug"`
	ImageURL    string   `josn:"image_url"`
}

func FormatFood(food Food) FoodFormatter {
	var foodFormat FoodFormatter

	foodFormat.ID = food.ID
	foodFormat.Name = food.Name
	foodFormat.Description = food.Description

	// apply ingredients to slice
	var ingredients []string
	for _, ingredient := range strings.Split(food.Ingredients, ",") {
		ingredients = append(ingredients, strings.TrimSpace(ingredient))
	}
	foodFormat.Ingredients = ingredients
	foodFormat.Slug = food.Slug

	// images URL
	if len(food.FoodImages) > 0 {
		foodFormat.ImageURL = food.FoodImages[0].Path
	}
	return foodFormat
}

func FormatFoods(foods []Food) []FoodFormatter {
	var foodsFormat []FoodFormatter

	for _, food := range foods {
		foodsFormat = append(foodsFormat, FormatFood(food))
	}

	return foodsFormat
}
