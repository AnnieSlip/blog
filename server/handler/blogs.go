package handler

import (
	"blogs/server/db"
	"context"
	"net/http"
	"strconv"

	"blogs/server/request"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func getBlogs(database *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		blogs, err := db.BlogList(context.Background(), database)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, blogs.Response())
	}

}

func addBlog(database *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r request.BlogCreate
		err := c.Bind(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		blog := r.DB()
		err = db.BlogAdd(context.Background(), database, blog)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}

}

func blogByID(database *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		blog, err := db.BlogByID(context.Background(), database, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if blog == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad blog id"})
			return

		}

		c.JSON(http.StatusOK, blog.Response())

	}
}
