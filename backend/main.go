package main

import (
	"log"

	_ "github.com/ezekielnizamani/JobScam/docs"
	v1 "github.com/ezekielnizamani/JobScam/internal/api/v1"
	"github.com/ezekielnizamani/JobScam/internal/database"
	"github.com/ezekielnizamani/JobScam/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Auto-migrate models
	if err := database.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Could not auto-migrate database: %v", err)
	}

	// Swagger
	r := gin.Default()

	v1.SetupRouter(r)

	// Serve Swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
