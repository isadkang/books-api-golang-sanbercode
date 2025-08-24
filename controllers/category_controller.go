package controllers

import (
	"books-api/config"
	"books-api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategories godoc
// @Summary Get all categories
// @Description Returns list of categories
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		rows.Scan(&cat.ID, &cat.Name)
		categories = append(categories, cat)
	}

	c.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary Create a category
// @Description Add a new category
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var req models.Category
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err := config.DB.Exec("INSERT INTO categories (name, created_at, created_by) VALUES ($1, now(), $2)", req.Name, "system")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created"})
}

// GetCategoryByID godoc
// @Summary Get category by ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
    id := c.Param("id")
    var cat models.Category
    err := config.DB.QueryRow("SELECT id, name FROM categories WHERE id=$1", id).Scan(&cat.ID, &cat.Name)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, cat)
}

// DeleteCategory godoc
// @Summary Delete a category
// @Tags categories
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
    id := c.Param("id")
    res, err := config.DB.Exec("DELETE FROM categories WHERE id=$1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    count, _ := res.RowsAffected()
    if count == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}

// GetBooksByCategory godoc
// @Summary Get books by category ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {array} models.Book
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /categories/{id}/books [get]
func GetBooksByCategory(c *gin.Context) {
    id := c.Param("id")
    rows, err := config.DB.Query("SELECT id, title, description, release_year, price, total_page, thickness FROM books WHERE category_id=$1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var books []models.Book
    for rows.Next() {
        var book models.Book
        rows.Scan(&book.ID, &book.Title, &book.Description, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness)
        books = append(books, book)
    }

    c.JSON(http.StatusOK, books)
}

