package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(g *gin.Engine, d *sqlx.DB) {
	g.GET("/hello", func(c *gin.Context) {
		// Respond with "Hello, World!" when the "/hello" route is accessed
		c.String(http.StatusOK, "Hello, World!")
	})

	g.GET("/blogs", getBlogs(d))
	g.GET("/blogs/:id", blogByID(d))
	g.POST("/blogs/add", addBlog(d))
	g.POST("sign-up", signUp(d))

}
