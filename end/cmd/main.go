package main

import (
	"end/internal"
	"end/middleware"
	"end/router"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CrossDomain)

	app := internal.NewApp()

	router.Routers(r, app)

	go func() {
		err := r.Run(":8080")
		if err != nil {
			panic(err)
		}
	}()

	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q
}
