package handlers

import (
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	t, err := template.ParseFiles("content/views/index.html")
	if err != nil {
		stdErr(c, err)
		return
	}
	data := gin.H{"title": "Klitch", "now": time.Now()}
	if err := t.Execute(c.Writer, data); err != nil {
		stdErr(c, err)
		return
	}
}
