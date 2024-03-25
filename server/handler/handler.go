package handler

import (
	"net/http"

	"blogs/server/middleware"

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
	g.POST("sign-up", signUp(d))
	g.POST("log-in", login(d))
	g.GET("/categories", getCategories(d))

	// Protected routes (require authentication)
	protectedRoutes := g.Group("/protected")
	protectedRoutes.Use(middleware.AuthenticationMiddleware())
	{
		protectedRoutes.POST("/blogs/add", addBlog(d))
		protectedRoutes.PUT("/blogs/:id/edit", editBlog(d))

	}

}
