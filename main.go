package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/book", bookHandler)

	//TODO Versioning
	v1 := router.Group("/v1")
	v1.GET("/", rootV1Handler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Randi Firmansyah",
		"bio":  "Hello world",
	})
}

func rootV1Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Randi Firmansyah",
		"bio":  "Hello world",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

type BookInput struct {
	Title    string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required,number"`
	Subtitle string      `json:"sub_title" binding:"required"`
}

func bookHandler(c *gin.Context) {
	var book BookInput
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
