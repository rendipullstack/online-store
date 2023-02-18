package models

import "time"

type TransactionHistory struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	ProductID  uint      `json:"product_id"`
	Product    Product   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	UserID     uint      `json:"user_id"`
	User       User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Quantity   int       `json:"quantity"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TransactionInput struct {
	ProductID int `json:"product_id" binding:"required" example:"1"`
	Quantity  int `json:"quantity" binding:"required,min=0" example:"1"`
}

type TransactionPostResponse struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type UserTransactionResponse struct {
	ID         uint            `json:"id"`
	ProductID  uint            `json:"product_id"`
	UserID     uint            `json:"user_id"`
	Quantity   int             `json:"quantity"`
	TotalPrice int             `json:"total_price"`
	Product    ProductResponse `json:"product"`
}

type TransactionResponse struct {
	ID         uint            `json:"id"`
	ProductID  uint            `json:"product_id"`
	UserID     uint            `json:"user_id"`
	Quantity   int             `json:"quantity"`
	TotalPrice int             `json:"total_price"`
	Product    ProductResponse `json:"product"`
	User       UserResponse    `json:"user"`
}
