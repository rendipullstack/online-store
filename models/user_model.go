package models

import "time"

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegisterInput struct {
	FullName string `json:"full_name" binding:"required" example:"hanif"`
	Email    string `json:"email" binding:"required,email" example:"me@hanifz.com"`
	Password string `json:"password" binding:"required,min=6" example:"qweqwe"`
}

type UserLoginInput struct {
	Email    string `json:"email" binding:"required,email" example:"me@hanifz.com"`
	Password string `json:"password" binding:"required,min=6" example:"qweqwe"`
}

type UserResponseRegister struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserResponseUpdate struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
