package handler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"go-foods/foods"
	"go-foods/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type foodsHandler struct {
	serviceFoods foods.Service
}

func NewFoodHandler(serviceFoods foods.Service) *foodsHandler {
	return &foodsHandler{serviceFoods}
}

// Create Food
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

// Create Image Food
func (h *foodsHandler) UploadImageFood(c *gin.Context) {
	var input foods.FoodImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error1", helper.FormatError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("img")
	if err != nil {
		errMsg := gin.H{
			"is_uploaded": false,
			"error":       helper.FormatError(err),
		}
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error2", errMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	fileName := []byte(file.Filename)
	fileNameHash := fmt.Sprintf("%x", md5.Sum(fileName))
	dirSaved := fmt.Sprintf("images/%s-%s-%s", time.Now().Format("20060102150405"), fileNameHash, file.Filename)

	err = c.SaveUploadedFile(file, dirSaved)

	if err != nil {
		errMsg := gin.H{
			"is_uploaded": false,
			"error":       helper.FormatError(err),
		}
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error3", errMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.serviceFoods.CreateImageFood(input, dirSaved)

	if err != nil {
		errMsg := gin.H{
			"is_uploaded": false,
			"error":       helper.FormatError(err),
		}
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error4", errMsg)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := helper.FormatResponse("success", http.StatusOK, "success create image food", data)
	c.JSON(http.StatusOK, response)

}

// Get All Foods
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

// Get Food by ID
func (h *foodsHandler) GetFood(c *gin.Context) {
	var input foods.FoodDetail

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusBadRequest, "error", helper.FormatError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	food, err := h.serviceFoods.GetFood(input.FoodID)

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error", helper.FormatError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if food.ID == 0 {
		response := helper.FormatResponse("failed", http.StatusNotFound, "failed", helper.FormatError(errors.New("food not found!")))
		c.JSON(http.StatusNotFound, response)
		return
	}

	formatter := foods.FormatFood(food)
	response := helper.FormatResponse("success", http.StatusOK, "success get detail food data", formatter)
	c.JSON(http.StatusOK, response)
}

// update food
func (h *foodsHandler) UpdateFood(c *gin.Context) {
	var input foods.FoodEdit

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusBadRequest, "error", helper.FormatError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	food, err := h.serviceFoods.UpdateFood(input)

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error4", helper.FormatError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.FormatResponse("success", http.StatusOK, "success get detail food data", foods.FormatFood(food))
	c.JSON(http.StatusOK, response)
}

// Delete foods
func (h *foodsHandler) DeleteFood(c *gin.Context) {
	var input foods.FoodDeelete

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error", helper.FormatError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.serviceFoods.DeleteFood(int(input.FoodID))

	if err != nil {
		response := helper.FormatResponse("failed", http.StatusUnprocessableEntity, "error", helper.FormatError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.FormatResponse("success", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
	return
}
