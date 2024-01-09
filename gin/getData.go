package main

import "github.com/gin-gonic/gin"

// get method
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hey gin",
	})
}
