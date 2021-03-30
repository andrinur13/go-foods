package foods

import "gorm.io/gorm"

type Repository interface {
	Save(food Food) (Food, error)
	GetAll() ([]Food, error)
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

	err := r.db.Find(&allFood).Error

	if err != nil {
		return allFood, err
	}

	return allFood, nil
}
