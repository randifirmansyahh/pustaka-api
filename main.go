package main

import (
	"log"
	"pustaka-golang/book"
	"pustaka-golang/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}

	db.AutoMigrate(&book.Book{})

	//TODO Repository
	bookRepository := book.NewRepository(db)

	// //TODO FindAll
	// books, err := bookRepository.FindAll()

	// for _, book := range books {
	// 	fmt.Println("Title:", book.Title)
	// }

	//TODO FindByID
	// book, err := bookRepository.FindByID(2)
	// fmt.Println("Title:", book.Title)

	//TODO BACA DOCUMENTASI GORM UNTUK TAU LEBIH BANYAK QUERY

	//TODO Create
	book := book.Book{
		Title:       "judul",
		Price:       3000,
		Description: "Good book",
		Discount:    1,
		Rating:      4,
	}

	bookRepository.Create(book)

	//create
	// book := book.Book{}
	// book.Title = "judul"
	// book.Price = 1000
	// book.Discount = 2
	// book.Rating = 5
	// book.Description = "deskripsi buku"

	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error creating book record")
	// 	fmt.Println("==========================")
	// }
	//create

	//read
	//one class/struct
	// var book book.Book

	//array
	//var books []book.Book

	//first data & just one
	//debug() for see the query selected
	// err = db.Debug().First(&book).Error

	//by id
	//err = db.Debug().First(&book, 1).Error

	//all data
	// err = db.Debug().Find(&books).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error read books record")
	// 	fmt.Println("==========================")
	// }

	//where
	// err = db.Debug().Where("price", 1000).Find(&books).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error read books record")
	// 	fmt.Println("==========================")
	// }

	//like
	// err = db.Debug().Where("title LIKE ?", "%j%").Find(&books).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error read books record")
	// 	fmt.Println("==========================")
	// }

	// //with loop
	// for _, b := range books {
	// 	fmt.Println("Title :", b.Title)
	// 	fmt.Println("Object :", b)
	// }
	//read

	// //Update
	// var book book.Book

	// err = db.Debug().Where("id", 1).First(&book).Error

	// book.Title = "Ganti Judul"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error update books record")
	// 	fmt.Println("==========================")
	// }
	// //Update

	//Delete
	// var book book.Book

	// err = db.Debug().Where("id", 1).First(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error find books record")
	// 	fmt.Println("==========================")
	// }

	// err = db.Delete(&book).Error

	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error Delete books record")
	// 	fmt.Println("==========================")
	// }
	//Delete

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
