package handler

import (
	"go-foods/foods"
	"go-foods/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type foodsHandler struct {
	serviceFoods foods.Service
}

func NewFoodHandler(serviceFoods foods.Service) *foodsHandler {
	return &foodsHandler{serviceFoods}
}

func (h *foodsHandler) CreateFood(c *gin.Context) {
	var input foods.FoodInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error", helper.FormatError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newFood, err := h.serviceFoods.CreateFood(input)

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error", helper.FormatError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.FormatResponse("success", http.StatusOK, "success create food data", newFood)
	c.JSON(http.StatusOK, response)
}

func (h *foodsHandler) GetAllFoods(c *gin.Context) {
	allFoods, err := h.serviceFoods.GetFoodAll()

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusBadRequest, "error", helper.FormatError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := foods.FormatFoods(allFoods)
	response := helper.FormatResponse("success", http.StatusOK, "success get all foods data", formatter)
	c.JSON(http.StatusOK, response)
}
