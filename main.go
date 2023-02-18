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

	_ "online-store/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Online Store Challenge API
// @version         1.0
// @description     Software Engineer Challenge PT. Synapsis Sinergi Digital.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Mochammad Hanif
// @contact.url    http://www.github.com/nifz
// @contact.email  ochammadhanif@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://online-store-hanif.up.railway.app
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	cartRepository := repositories.NewCartRepository(db)
	cartService := services.NewCartService(cartRepository, productRepository, userRepository, categoryRepository)
	cartController := controllers.NewCartController(cartService)

	transactionRepository := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository, productRepository, userRepository, categoryRepository)
	transactionController := controllers.NewTransactionController(transactionService)

	router := gin.Default()
	routes.UserRoute(router, userController)
	routes.CategoryRoutes(router, categoryController)
	routes.ProductRoutes(router, productController)
	routes.CartRoute(router, cartController)
	routes.TransactionRoute(router, transactionController)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + port)
}
