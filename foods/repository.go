package foods

import "gorm.io/gorm"

type Repository interface {
	Save(food Food) (Food, error)
	GetByID(foodID int) (Food, error)
	GetAll() ([]Food, error)
	DeleteFood(foodID int) error
	Update(food Food) (Food, error)
	SaveImage(foodImage FoodImage) (FoodImage, error)
	MarkAllImageIsNonPrimary(foodID int) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(food Food) (Food, error) {
	err := r.db.Save(&food).Error

	if err != nil {
		return food, err
	}

	return food, nil
}

func (r *repository) GetAll() ([]Food, error) {
	var allFood []Food

	err := r.db.Preload("FoodImages", "is_primary = 1").Find(&allFood).Error

	if err != nil {
		return allFood, err
	}

	return allFood, nil
}

func (r *repository) SaveImage(foodImage FoodImage) (FoodImage, error) {
	err := r.db.Save(&foodImage).Error

	if err != nil {
		return foodImage, err
	}

	return foodImage, nil
}

func (r *repository) MarkAllImageIsNonPrimary(foodID int) (bool, error) {
	err := r.db.Model(&FoodImage{}).Where("food_id = ?", foodID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

// Get food by id
func (r *repository) GetByID(foodID int) (Food, error) {
	food := Food{}

	err := r.db.Where("id = ?", foodID).Find(&food).Error

	if err != nil {
		return food, err
	}

	return food, nil
}

// Delete
func (r *repository) DeleteFood(foodID int) error {
	foodDeleted := Food{}
	err := r.db.Where("id = ?", foodID).Delete(foodDeleted).Error

	if err != nil {
		return err
	}

	return nil
}

// update food
func (r *repository) Update(food Food) (Food, error) {
	err := r.db.Where("id = ?", int(food.ID)).Save(&food).Error

	if err != nil {
		return food, err
	}

	return food, nil
}
