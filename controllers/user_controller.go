package controllers

import (
    "net/http"
    "go_crud/database"
    "go_crud/models"

    "github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, gin.H{"data": users})
}

// Get single user by ID
func GetUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// Create user (validated at route level)
func CreateUser(c *gin.Context) {
    input := c.MustGet("validatedBody").(*models.CreateUserInput)

    user := models.User{
        Name:  input.Name,
        Email: input.Email,
        Age:   input.Age,
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": user})
}

// Update user (partial update, validated at route level)
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    input := c.MustGet("validatedBody").(*models.UpdateUserInput)

    updates := map[string]interface{}{}
    if input.Name != nil {
        updates["name"] = *input.Name
    }
    if input.Email != nil {
        updates["email"] = *input.Email
    }
    if input.Age != nil {
        updates["age"] = *input.Age
    }

    if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

// Delete user
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    database.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
