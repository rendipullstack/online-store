package repositories

import (
	"online-store/helpers"
	"online-store/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(cart models.Cart) (models.Cart, error)
	GetCarts(userID uint) ([]models.Cart, error)
	GetDataByID(cartID uint) (models.Cart, error)
	DeleteCart(cart models.Cart) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Preload("Product").Preload("User").Create(&cart).Error
	return cart, helpers.ReturnIfError(err)
}

func (r *cartRepository) GetCarts(userID uint) ([]models.Cart, error) {
	var (
		carts []models.Cart
	)

	db := r.db
	if userID > 0 {
		db = db.Where("user_id = ?", userID)
	}

	err := db.Find(&carts).Preload("Product").Preload("User").Find(&carts).Error

	return carts, helpers.ReturnIfError(err)
}

func (r *cartRepository) DeleteCart(cart models.Cart) error {
	err := r.db.Delete(&cart).Error
	return helpers.ReturnIfError(err)
}

func (r *cartRepository) GetDataByID(cartID uint) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Where("id = ?", cartID).First(&cart).Error
	return cart, helpers.ReturnIfError(err)
}
