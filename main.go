package main

import (
	"gin-boiler-plate/db"
	"gin-boiler-plate/middleware"
	"gin-boiler-plate/router"
	"github.com/gin-gonic/gin"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(0)

	db.Init()

	app := gin.Default()

	router.SetRoutes(app)
	middleware.Set(app)

	app.Run(":5500")
}
