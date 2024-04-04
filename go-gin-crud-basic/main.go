package main

// @title User CRUD API
// @version 1.0
// @description API for managing users with CRUD operations
// @host localhost:8080
// @BasePath /api/v1
// @schemes http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @name User
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{Id: 1, Name: "Jon Doe", Age: 28},
	{Id: 2, Name: "Jane Doe", Age: 25},
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample service for managing users.
func main() {
	// Create a new Gin router
	router := gin.Default()

	// Serve Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Define routes for CRUD operations
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	// Run the HTTP server on port 8080
	router.Run(":8080")
}

// @Summary Get all users
// @Description Get a list of all users
// @Produce json
// @Success 200 {object} []User
// @Router /users [get]
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// @Summary Get user by ID
// @Description Get a user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	// Convert id to int
	// In a real application, you would handle errors appropriately
	userID := parseInt(id)

	for _, user := range users {
		if user.Id == userID {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// @Summary Create a new user
// @Description Create a new user
// @Accept json
// @Produce json
// @Param user body User true "User object"
// @Success 201 {object} User
// @Router /users [post]
func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real application, you might want to generate a unique ID
	newUser.Id = len(users) + 1
	users = append(users, newUser)

	c.JSON(http.StatusCreated, newUser)
}

// @Summary Update user by ID
// @Description Update a user by ID
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "Updated user object"
// @Success 200 {object} User
// @Router /users/{id} [put]
func updateUser(c *gin.Context) {
	id := c.Param("id")
	// Convert id to int
	userID := parseInt(id)

	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, user := range users {
		if user.Id == userID {
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// @Summary Delete user by ID
// @Description Delete a user by ID
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	// Convert id to int
	userID := parseInt(id)

	for i, user := range users {
		if user.Id == userID {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// parseInt converts a string to an integer
func parseInt(s string) int {
	// In a real application, you would handle errors appropriately
	result, _ := strconv.Atoi(s)
	return result
}
