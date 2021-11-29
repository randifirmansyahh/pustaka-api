package handler

import (
	"fmt"
	"net/http"
	"pustaka-golang/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Randi Firmansyah",
		"bio":  "Hello world",
	})
}

func RootV1Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Randi Firmansyah",
		"bio":  "Hello world",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

func BookHandler(c *gin.Context) {
	var book book.BookInput
	err := c.ShouldBindJSON(&book)

	if err != nil {

		errorMassages := []string{}

		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     book.Title,
		"price":     book.Price,
		"sub_title": book.Subtitle,
	})
}
