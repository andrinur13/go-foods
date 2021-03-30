package foods

import "strings"

type Service interface {
	CreateFood(input FoodInput) (Food, error)
	GetFoodAll() ([]Food, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateFood(input FoodInput) (Food, error) {

	food := Food{
		Name:        input.Name,
		Description: input.Description,
		Ingredients: input.Ingredients,
	}
	food.Slug = strings.ReplaceAll(input.Name, " ", "-")
	food.Slug = strings.ToLower(food.Slug)

	newFood, err := s.repository.Save(food)

	if err != nil {
		return newFood, err
	}

	return newFood, nil
}

func (s *service) GetFoodAll() ([]Food, error) {
	allFoods, err := s.repository.GetAll()

	if err != nil {
		return allFoods, err
	}

	return allFoods, nil
}
