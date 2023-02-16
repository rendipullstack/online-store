package main

import (
	"log"
	"online-store/configs"
	"online-store/controllers"
	"online-store/repositories"
	"online-store/routes"
	"online-store/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	dbUsername := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")

	db, err := configs.GetDatabase(dbUsername, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	categoryRepository := repositories.NewCategoryRepositories(db)
	categoryService := services.NewCategoryServices(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)

	productRepository := repositories.NewProductRepositories(db)
	productService := services.NewProductServices(productRepository, categoryRepository)
	productController := controllers.NewProductController(productService)

	transactionRepository := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository, productRepository, userRepository, categoryRepository)
	transactionController := controllers.NewTransactionController(transactionService)

	router := gin.Default()
	routes.UserRoute(router, userController)
	routes.CategoryRoutes(router, categoryController)
	routes.ProductRoutes(router, productController)
	routes.TransactionRoute(router, transactionController)

	router.Run(":" + port)
}
