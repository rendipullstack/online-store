package services

import (
	"errors"
	"online-store/helpers"
	"online-store/models"
	"online-store/repositories"
)

type TransactionService interface {
	CreateTransaction(input models.TransactionInput, userID uint) (models.TransactionPostResponse, error)
	GetTransactions(userID uint) ([]models.UserTransactionResponse, error)
	GetAllTransactions() ([]models.TransactionResponse, error)
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
	productRepository     repositories.ProductRepositories
	userRepository        repositories.UserRepository
	categoryRepository    repositories.CategoryRepositories
}

func NewTransactionService(
	transactionRepository repositories.TransactionRepository,
	productRepository repositories.ProductRepositories,
	userRepository repositories.UserRepository,
	categoryRepository repositories.CategoryRepositories,
) *transactionService {
	return &transactionService{
		transactionRepository,
		productRepository,
		userRepository,
		categoryRepository,
	}
}

func (s *transactionService) CreateTransaction(input models.TransactionInput, userID uint) (models.TransactionPostResponse, error) {
	var (
		transactionResponse models.TransactionPostResponse
		transaction         models.TransactionHistory
	)

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return transactionResponse, err
	}
	product, err := s.productRepository.GetDataByID(uint(input.ProductID))
	if err != nil {
		return transactionResponse, err
	}
	category, err := s.categoryRepository.GetDataByID(product.CategoryID)
	if err != nil {
		return transactionResponse, err
	}

	if *product.Stock < input.Quantity {
		return transactionResponse, errors.New("product is not available")
	}

	totalPrice := product.Price * input.Quantity
	category, err = s.categoryRepository.UpdateCategory(category)
	if err != nil {
		return transactionResponse, err
	}

	stock := *product.Stock - input.Quantity
	*product.Stock = stock
	product, err = s.productRepository.UpdateProduct(product)
	if err != nil {
		return transactionResponse, err
	}

	transaction.UserID = user.ID
	transaction.ProductID = product.ID
	transaction.Quantity = input.Quantity
	transaction.TotalPrice = totalPrice

	transaction, err = s.transactionRepository.CreateTransaction(transaction)
	transactionResponse = models.TransactionPostResponse{
		TotalPrice:   transaction.TotalPrice,
		Quantity:     transaction.Quantity,
		ProductTitle: product.Title,
	}

	return transactionResponse, helpers.ReturnIfError(err)
}

func (s *transactionService) GetTransactions(userID uint) ([]models.UserTransactionResponse, error) {
	transactions, err := s.transactionRepository.GetTransactions(userID)
	var transactionResponses []models.UserTransactionResponse
	for _, transaction := range transactions {
		transactionResponse := models.UserTransactionResponse{
			ID:         transaction.ID,
			ProductID:  transaction.ProductID,
			UserID:     transaction.UserID,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: models.ProductResponse{
				ID:         transaction.Product.ID,
				Title:      transaction.Product.Title,
				Price:      transaction.Product.Price,
				Stock:      *transaction.Product.Stock,
				CategoryID: transaction.Product.Category.ID,
				CreatedAt:  transaction.Product.CreatedAt,
			},
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return transactionResponses, helpers.ReturnIfError(err)
}

func (s *transactionService) GetAllTransactions() ([]models.TransactionResponse, error) {
	transactions, err := s.transactionRepository.GetTransactions(0)

	var (
		transactionResponses []models.TransactionResponse
	)

	for _, transaction := range transactions {
		transactionResponse := models.TransactionResponse{
			ID:         transaction.ID,
			ProductID:  transaction.ProductID,
			UserID:     transaction.UserID,
			Quantity:   transaction.Quantity,
			TotalPrice: transaction.TotalPrice,
			Product: models.ProductResponse{
				ID:         transaction.Product.ID,
				Title:      transaction.Product.Title,
				Price:      transaction.Product.Price,
				Stock:      *transaction.Product.Stock,
				CategoryID: transaction.Product.Category.ID,
				CreatedAt:  transaction.Product.CreatedAt,
			},
			User: models.UserResponse{
				ID:        transaction.User.ID,
				Email:     transaction.User.Email,
				FullName:  transaction.User.FullName,
				CreatedAt: transaction.User.CreatedAt,
				UpdatedAt: transaction.User.UpdatedAt,
			},
		}
		transactionResponses = append(transactionResponses, transactionResponse)
	}

	return transactionResponses, helpers.ReturnIfError(err)
}
