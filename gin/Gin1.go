package main

import (
	"awesomeProject/logs"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	//creates gin router with default middleware and logger , we can also do that explicitly
	router := gin.Default()

	logrus.Println("logrus printing.......")
	//router := gin.New()

	gin.ForceConsoleColor()

	/*router := gin.New()
	router.Use(gin.Logger())*/

	//These middleware will be applied to all the endpoints
	/*router.Use(middleware.AUthenticate, middleware.AddHeader)
	router.Use(gin.LoggerWithFormatter(logs.FormatLogs))*/

	//we can create basic authentication as given below
	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass",
	})

	//the context contains all the information about request that are coming and response that are returning
	router.POST("/getDataPost", getDataPost)
	router.GET("/getPathData/:name/:age", getPathData)

	//calling the function
	logs.Log()

	// : is important before the port number
	//http.ListenAndServe(":9090", router)

	//we can create routing groups just like controllers in  java like this
	//also i have passed the basic authentication here i.e username and password

	//we can use middleware for particular group as well
	admin := router.Group("/admin", middleware.AUthenticate)
	{
		admin.GET("/getData", getData)
	}

	client := router.Group("/client", auth)
	{
		client.GET("/getQuery", getQueryData)

	}

	//we can pass custom http configuration to the server
	/*server := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()*/

	//below given line internally calls ListenAndAServe() method
	router.Run()
}
