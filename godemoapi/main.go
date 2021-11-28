package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		//fmt.Printf("error while connectr sqllite %v", err.Error())
		panic("Cannot Connect db")
	}

	db.AutoMigrate(&Book{})

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "pong",
		})
	})
	r.POST("/books", NewBook)
	r.GET("/books", ListBook)
	r.GET("/books/:id", GetBook)
	r.Run()

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"Message": "pong",
	// 	})
	// })

	// r.Run()
}

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
}

func NewBook(c *gin.Context) {
	var book Book
	if err := c.Bind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func ListBook(c *gin.Context) {

	var books []Book

	result := db.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)

}

func GetBook(c *gin.Context) {

	id := c.Param("id")

	n, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var book Book
	result := db.Find(&book, n)
	if err = result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, book)
}
