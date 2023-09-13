package main

import (
	"html/template"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		t, err := template.ParseFiles("content/views/index.html")
		if err != nil {
			stdErr(c, err)
			return
		}
		data := map[string]any{"title": "Klitch", "now": time.Now()}
		if err := t.Execute(c.Writer, data); err != nil {
			stdErr(c, err)
		}
	})

	g.GET("/htmx/time", func(c *gin.Context) {
		t, err := template.ParseFiles("content/views/htmx/time.html")
		if err != nil {
			stdErr(c, err)
			return
		}
		data := map[string]any{"now": time.Now()}
		if err := t.Execute(c.Writer, data); err != nil {
			stdErr(c, err)
			return
		}
	})

	log.Fatalln(g.Run(":8080"))

}

func stdErr(c *gin.Context, err error) {
	c.JSON(500, gin.H{"path": c.Request.URL.Path, "error": err.Error()})
}
