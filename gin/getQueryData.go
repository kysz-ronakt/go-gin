package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func getQueryData(c *gin.Context) {

	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{

		/*"name": c.GetQuery(name),
		"age":  c.GetQuery(age),*/
		"name": name,
		"age":  age,
	})

	fmt.Println("Got the user info... " + "name: " + name + " age: " + age)
}
