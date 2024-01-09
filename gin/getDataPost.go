package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

//creating struct to return response and err

type DataAndError struct {
	Value []byte
	Err   error
}

// post method with requestBody handling
func getDataPost(c *gin.Context) {

	body := c.Request.Body

	value, err := ioutil.ReadAll(body)

	dataWithError := DataAndError{
		Value: value,
		Err:   err,
	}

	c.JSON(200, gin.H{
		"data":     "Hey gin i'm post method",
		"bodyData": string(dataWithError.Value),
		"error":    dataWithError.Err,
	})
}
