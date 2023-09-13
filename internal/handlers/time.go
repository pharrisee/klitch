package handlers

import (
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

func Time(c *gin.Context) {
	t, err := template.ParseFiles("content/views/htmx/time.html")
	if err != nil {
		stdErr(c, err)
		return
	}
	data := gin.H{"now": time.Now()}
	if err := t.Execute(c.Writer, data); err != nil {
		stdErr(c, err)
		return
	}
}
