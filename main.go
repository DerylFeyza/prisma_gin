package main

import (
	"context"
	"log"

	"github.com/DerylFeyza/prisma-gonic/prisma/db"
	"github.com/DerylFeyza/prisma-gonic/services/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := db.NewClient()
	if err := client.Connect(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer func() {
		if err := client.Disconnect(); err != nil {
			log.Fatalf("Failed to disconnect from the database: %v", err)
		}
	}()
	r := gin.Default()
	ctx := context.Background()

	r.POST("/user", func(c *gin.Context) {
		controllers.CreateUser(ctx, c, client)
	})

	r.Run()
}
