package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {

	c.JSON(200, gin.H{
		"value": "ok",
	})
}