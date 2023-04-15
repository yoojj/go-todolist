package main

import (
	"net/http"
	"os"

	"todo-list/common"
	"todo-list/router"
)

func init() {
	common.ConfigInit()
}

func main() {

	routerInit := router.RouterInit()

	port := os.Getenv("PORT")
	server := &http.Server{
		Addr:    ":" + port,
		Handler: routerInit,
	}

	server.ListenAndServe()

}
