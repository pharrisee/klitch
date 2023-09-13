package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pharrisee/klitch/internal/handlers"
)

func main() {
	g := gin.Default()

	g.GET("/", handlers.Index)

	g.GET("/htmx/time", handlers.Time)

	log.Fatalln(g.Run(":8080"))

}
