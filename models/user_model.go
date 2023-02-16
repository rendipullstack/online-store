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
