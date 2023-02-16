package models

import "time"

type Product struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      *int      `json:"stock"`
	CategoryID uint      `json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required,min=0,max=50000000"`
	Stock      int    `json:"stock" binding:"required,gte=5"`
	CategoryID uint   `json:"category_id"`
}

type ProductResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type ProductResponseData struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Title      string    `json:"title"`
	Price      string    `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProductResponseUpdate struct {
	Product ProductResponseData `json:"product"`
}
