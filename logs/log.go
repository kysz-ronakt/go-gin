package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func Log() {

	//this will check for the file , if there is no then it'll create new one
	f, _ := os.Create("ginLogging.log")

	//default IO writer
	gin.DefaultWriter = io.MultiWriter(f)

}

// define logger middleware
func FormatLogs(param gin.LogFormatterParams) string {
	return fmt.Sprintf("{%s - [%s] \"%s %s %s %d %s \"%s\" %s\"} \n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
