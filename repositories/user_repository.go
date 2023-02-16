package repositories

import (
	"online-store/helpers"
	"online-store/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(userID uint) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user models.User) (models.User, error) {
	err := ur.db.Create(&user).Error
	return user, helpers.ReturnIfError(err)
}

func (ur *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := ur.db.Where("email = ?", email).Find(&user).Error
	return user, helpers.ReturnIfError(err)
}

func (ur *userRepository) GetUserByID(userID uint) (models.User, error) {
	var user models.User
	err := ur.db.Where("id = ?", userID).First(&user).Error
	return user, helpers.ReturnIfError(err)
}
