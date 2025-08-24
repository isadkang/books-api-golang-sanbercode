package controllers

import (
	"books-api/config"
	"books-api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary Get all books
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /books [get]
func GetBooks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, description, release_year, price, total_page, thickness, category_id FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		rows.Scan(&book.ID, &book.Title, &book.Description, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID)
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

// CreateBook godoc
// @Summary Create a book
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var req models.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.ReleaseYear < 1980 || req.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "release_year must be between 1980 and 2024",
		})
		return
	}

	if req.TotalPage <= 100 {
		req.Thickness = "tipis"
	} else if req.TotalPage <= 200 {
		req.Thickness = "sedang"
	} else {
		req.Thickness = "tebal"
	}

	_, err := config.DB.Exec(`
		INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,now(),$9)`,
		req.Title, req.Description, req.ImageURL, req.ReleaseYear, req.Price, req.TotalPage, req.Thickness, req.CategoryID, "system")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created"})
}

// GetBookByID godoc
// @Summary Get book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /books/{id} [get]
func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    err := config.DB.QueryRow("SELECT id, title, description, release_year, price, total_page, thickness FROM books WHERE id=$1", id).
        Scan(&book.ID, &book.Title, &book.Description, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness)
    if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete a book
// @Tags books
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    res, err := config.DB.Exec("DELETE FROM books WHERE id=$1", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    count, _ := res.RowsAffected()
    if count == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}