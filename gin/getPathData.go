package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPathData(ctx *gin.Context) {

	name := ctx.Param("name")
	age := ctx.Param("age")

	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})

	fmt.Println("Got the user info... " + "name: " + name + " age: " + age)
}
