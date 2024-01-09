package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	//creates gin router with default middleware and logger , we can also do that explicitly
	router := gin.Default()

	/*router := gin.New()
	router.Use(gin.Logger())*/

	//the context contains all the information about request that are coming and response that are returning
	router.GET("/getData", getData)
	router.POST("/getDataPost", getDataPost)
	router.GET("/getQuery", getQueryData)
	router.GET("/getPathData/:name/:age", getPathData)

	router.Run()
}
