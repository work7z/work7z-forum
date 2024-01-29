package handlers

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	// root
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/v2")
	})
}
