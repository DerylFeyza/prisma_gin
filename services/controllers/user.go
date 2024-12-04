package controllers

import (
	"context"
	"net/http"

	"github.com/DerylFeyza/prisma-gonic/prisma/db"
	skibi "github.com/DerylFeyza/prisma-gonic/thirdparty/bcrypt"
	"github.com/gin-gonic/gin"
)

// CreateUser handles the creation of a new user.
func CreateUser(ctx context.Context, c *gin.Context, client *db.PrismaClient) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := skibi.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	created, err := client.User.CreateOne(
		db.User.Username.Set(req.Username),
		db.User.Email.Set(req.Email),
		db.User.Password.Set(hashedPassword),
	).Exec(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    created,
		"message": "User created successfully",
	})
}

func GetAllUser(ctx context.Context, c *gin.Context, client *db.PrismaClient) {

	users, err := client.User.FindMany().Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
