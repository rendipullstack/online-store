package models

import "time"

type Category struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Type      string    `json:"type"`
	Product   []Product `json:"products"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryInput struct {
	Type string `json:"type" binding:"required" example:"t-shirt"`
}

type CategoryResponsePost struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryResponseGet struct {
	ID        uint              `gorm:"primarykey" json:"id"`
	Type      string            `json:"type"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Product   []ProductResponse `json:"products"`
}

type CategoryResponsePatch struct {
	ID        uint      `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}
