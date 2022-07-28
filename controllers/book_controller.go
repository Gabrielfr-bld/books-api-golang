package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Gabrielfr-bld/books-api-golang/models"
	"github.com/Gabrielfr-bld/books-api-golang/database"
)

func ShowBook(c *gin.Context) {

	c.JSON(200, gin.H{
		"value": "ok",
	})
}