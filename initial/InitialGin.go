package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// command to install gin from terminal
// go get -u github.com/gin-gonic/gin
func main() {

	//it will create a gin router for
	router := gin.Default()

	//loading all the templates
	router.LoadHTMLGlob("initial/templates/*")

	//defining the router
	router.GET("/", func(c *gin.Context) {

		//call the html method to render template
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "home page",
			},
		)
	})

	router.Run()
}
