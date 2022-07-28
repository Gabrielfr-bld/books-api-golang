package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Gabrielfr-bld/books-api-golang/models"
	"github.com/Gabrielfr-bld/books-api-golang/database"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer",
		})

		return 
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Book not found: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})

		return
	}

	c.JSON(201, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": "not found list: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})

		return
	}

	err = db.Save(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.C) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Id has to be integer",
		})

		return 
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})

		return
	}

	c.Status(204)
}