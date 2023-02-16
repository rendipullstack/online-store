package repositories

import (
	"online-store/helpers"
	"online-store/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepositories interface {
	PostCategory(category models.Category) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	GetDataByID(categoryID uint) (models.Category, error)
	DeleteCategory(category models.Category) error
}

type categoryRepositories struct {
	db *gorm.DB
}

func NewCategoryRepositories(db *gorm.DB) *categoryRepositories {
	return &categoryRepositories{db}
}

func (cr *categoryRepositories) PostCategory(category models.Category) (models.Category, error) {
	err := cr.db.Create(&category).Error
	return category, helpers.ReturnIfError(err)
}

func (cr *categoryRepositories) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := cr.db.Preload(clause.Associations).Find(&categories).Error

	return categories, helpers.ReturnIfError(err)
}

func (cr *categoryRepositories) UpdateCategory(category models.Category) (models.Category, error) {
	err := cr.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&category).Error
	return category, helpers.ReturnIfError(err)
}

func (cr *categoryRepositories) GetDataByID(categoryID uint) (models.Category, error) {
	var category models.Category
	err := cr.db.Where("id = ?", categoryID).First(&category).Error
	return category, helpers.ReturnIfError(err)
}

func (cr *categoryRepositories) DeleteCategory(category models.Category) error {
	err := cr.db.Delete(&category).Error

	return helpers.ReturnIfError(err)

}
