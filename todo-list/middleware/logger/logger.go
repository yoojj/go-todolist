package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Log struct {
	formatted string
	method    string
	path      string
	ip        string
}

func LoggingHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		t := time.Now()
		formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())

		c.Next()

		// c.Writer.Status()
		// c.Request.UserAgent()
		// c.Request.Referer()

		l := Log{
			formatted: formatted,
			method:    c.Request.Method,
			path:      c.Request.URL.Path,
			ip:        c.ClientIP(),
		}

		fmt.Println(l)

	}

}
