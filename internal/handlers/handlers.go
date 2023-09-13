package handlers

import "github.com/gin-gonic/gin"

func stdErr(c *gin.Context, err error) {
	c.JSON(500, gin.H{"path": c.Request.URL.Path, "error": err.Error()})
}
