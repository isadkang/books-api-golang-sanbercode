// @title Books API
// @version 1.0
// @description API for managing books and categories
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
package routes

import (
	"books-api/controllers"
	"books-api/middlewares"

	"github.com/gin-gonic/gin"

	_ "books-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running 🚀"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/api/users/login", controllers.Login)
	r.POST("/api/users/register", controllers.Register)

	category := r.Group("/api/categories")
	category.Use(middlewares.JWTAuthMiddleware())
	{
		category.GET("", controllers.GetCategories)
		category.POST("", controllers.CreateCategory)
		category.GET("/:id", controllers.GetCategoryByID)
		category.DELETE("/:id", controllers.DeleteCategory)
		category.GET("/:id/books", controllers.GetBooksByCategory)
	}

	book := r.Group("/api/books")
	book.Use(middlewares.JWTAuthMiddleware())
	{
		book.GET("", controllers.GetBooks)
		book.POST("", controllers.CreateBook)
		book.GET("/:id", controllers.GetBookByID)
		book.DELETE("/:id", controllers.DeleteBook)
	}

	return r
}
