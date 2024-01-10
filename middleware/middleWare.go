package middleware

import "github.com/gin-gonic/gin"

//middleware are the function that have access to request object and response object as well

//two types of syntax to write MW

func AUthenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token") == "Auth") {

		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Token not present",
			"Error":   c.Err(),
		})
		return
	}
	c.Next()
}

// another way to write middleware (Realtime use)
/*func AUthenticate() gin.HandlerFunc {

	return func(c *gin.Context) {
		if !(c.Request.Header.Get("Token") == "Auth") {

			c.AbortWithStatusJSON(500, gin.H{
				"Message": "Token not present",
				"Error":   c.Err(),
			})
			return
		}
		c.Next()
	}
}*/

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("key", "value")
	c.Next()
}
