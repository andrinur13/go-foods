package foods

import "strings"

type Service interface {
	CreateFood(input FoodInput) (Food, error)
	GetFoodAll() ([]Food, error)
	GetFood(foodID int) (Food, error)
	DeleteFood(foodID int) error
	CreateImageFood(input FoodImageInput, path string) (FoodImage, error)
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

func (s *service) GetFood(foodID int) (Food, error) {
	food, err := s.repository.GetByID(foodID)

	if err != nil {
		return food, err
	}

	return food, nil
}

func (s *service) CreateImageFood(input FoodImageInput, path string) (FoodImage, error) {
	foodImg := FoodImage{
		FoodID:    input.FoodID,
		Path:      path,
		IsPrimary: false,
	}

	if input.IsPrimary {
		foodImg.IsPrimary = true

		_, err := s.repository.MarkAllImageIsNonPrimary(input.FoodID)

		if err != nil {
			return FoodImage{}, err
		}
	}

	newImg, err := s.repository.SaveImage(foodImg)

	if err != nil {
		return newImg, err
	}

	return newImg, nil
}

func (s *service) DeleteFood(foodID int) error {
	err := s.repository.DeleteFood(foodID)

	if err != nil {
		return err
	}

	return nil
}
