package main

import (
	"fmt"
	"log"
	"pustaka-golang/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}

	fmt.Println("Database Connected")

	router := gin.Default()

	router.GET("/", handler.RootHandler)
	router.GET("/books/:id/:title", handler.BooksHandler)
	router.GET("/query", handler.QueryHandler)
	router.POST("/book", handler.BookHandler)

	//TODO Versioning
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootV1Handler)

	router.Run()
}
