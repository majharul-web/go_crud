package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name  string `json:"name" binding:"required,min=2,max=100"`
    Email string `json:"email" binding:"required,email" gorm:"unique"`
    Age   int    `json:"age" binding:"required,gte=0,lte=120"`
}
