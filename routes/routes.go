package routes

import (
	"books-api/controllers"
	"books-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running 🚀"})
	})

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
