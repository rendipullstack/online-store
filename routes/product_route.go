package routes

import (
	"online-store/controllers"
	"online-store/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, controller controllers.ProductController) {
	router.POST("/products", middlewares.AuthMiddleware, middlewares.AdminMiddleware, controller.CreateProduct)
	router.GET("/products", middlewares.AuthMiddleware, controller.GetProducts)
	router.DELETE("/products/:id", middlewares.AuthMiddleware, middlewares.AdminMiddleware, controller.DeleteProduct)
	router.PUT("/products/:id", middlewares.AuthMiddleware, middlewares.AdminMiddleware, controller.UpdateProduct)
}
