package models

import "gorm.io/gorm"

// User database model
type User struct {
    gorm.Model
    Name  string `json:"name"`
    Email string `json:"email" gorm:"unique"`
    Age   int    `json:"age"`
}

// DTO for creating a user
type CreateUserInput struct {
    Name  string `json:"name" binding:"required,min=2,max=100"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"required,gte=0,lte=120"`
}

// DTO for updating a user (PATCH - partial update)
type UpdateUserInput struct {
    Name  *string `json:"name" binding:"omitempty,min=2,max=100"`
    Email *string `json:"email" binding:"omitempty,email"`
    Age   *int    `json:"age" binding:"omitempty,gte=0,lte=120"`
}
