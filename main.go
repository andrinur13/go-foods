package main

import (
	"fmt"
	"go-foods/foods"
	"go-foods/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "andri:root@tcp(127.0.0.1:3306)/go_foods?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("connected to database")

	router := gin.Default()
	router.Static("/images", "./images")
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	// declare repo, serv, handler
	foodsRepository := foods.NewRepository(db)
	foodsService := foods.NewService(foodsRepository)
	foodsHandler := handler.NewFoodHandler(foodsService)

	api.POST("/foods", foodsHandler.CreateFood)
	api.GET("/foods", foodsHandler.GetAllFoods)
	api.POST("/food-image", foodsHandler.UploadImageFood)

	router.Run()
}
