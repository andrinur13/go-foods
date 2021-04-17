package main

import (
	"fmt"
	"go-foods/foods"
	"go-foods/handler"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get env
	userName := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	hostIP := os.Getenv("HOST")
	portServer := os.Getenv("PORTDB")
	dbName := os.Getenv("DATABASE")

	// read db
	dsnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, hostIP, portServer, dbName)

	// cadangan siapa tau gagal :v
	// dsn := "andri:root@tcp(127.0.0.1:3306)/go_foods?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := dsnString
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
	api.GET("/food/:id", foodsHandler.GetFood)
	api.PUT("/food", foodsHandler.UpdateFood)
	api.DELETE("/foods", foodsHandler.DeleteFood)

	api.POST("/food-image", foodsHandler.UploadImageFood)

	router.Run()
}
