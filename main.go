package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", rootHandler)
	route.GET("/books/:id/:title", booksHandler)
	route.GET("/query", queryHandler)
	route.POST("/book", bookHandler)

	route.Run()
}

func rootHandler(c *gin.Context) {
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
	Title    string
	Price    int
	Subtitle string `json:"sub_title"`
}

func bookHandler(c *gin.Context) {
	var book BookInput
	err := c.ShouldBindJSON(&book)

	if err != nil {
		log.Fatal()
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     book.Title,
		"price":     book.Price,
		"sub_title": book.Subtitle,
	})
}
