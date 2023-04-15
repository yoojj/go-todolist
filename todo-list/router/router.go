package router

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"todo-list/mvc/controller/task"
	"todo-list/mvc/controller/user"
)

func RouterInit() *gin.Engine {

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(gin.Logger())
	//r.Use(logger.LoggingHandler())

	baseUrl := os.Getenv("BASE_URL")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{baseUrl},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Delims("[[", "]]")

	r.LoadHTMLGlob("mvc/view/**/**/*")
	r.Static("/static", "mvc/view/static/")
	r.StaticFile("/favicon.ico", "mvc/view/static/images/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "index",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "404",
		})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    "METHOD_NOT_ALLOWED",
			"message": "405 method not allowed",
		})
	})

	task.Routes(r)
	user.Routes(r)

	return r
}
