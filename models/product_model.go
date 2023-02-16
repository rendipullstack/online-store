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
