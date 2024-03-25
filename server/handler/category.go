package handler

import (
	"blogs/server/db"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func getCategories(database *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cs, err := db.CaregoryList(context.Background(), database)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, cs.Response())
	}

}
