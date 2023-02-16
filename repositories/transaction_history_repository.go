package repositories

import (
	"online-store/helpers"
	"online-store/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.TransactionHistory) (models.TransactionHistory, error)
	GetTransactions(userID uint) ([]models.TransactionHistory, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(transaction models.TransactionHistory) (models.TransactionHistory, error) {
	err := r.db.Preload("Product").Preload("User").Create(&transaction).Error
	return transaction, helpers.ReturnIfError(err)
}

func (r *transactionRepository) GetTransactions(userID uint) ([]models.TransactionHistory, error) {
	var (
		transactions []models.TransactionHistory
	)

	db := r.db
	if userID > 0 {
		db = db.Where("user_id = ?", userID)
	}

	err := db.Find(&transactions).Preload("Product").Preload("User").Find(&transactions).Error

	return transactions, helpers.ReturnIfError(err)
}
