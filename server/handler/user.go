package handler

import (
	"blogs/server/domain"
	"blogs/server/request"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func signUp(database *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r request.UserCreate
		err := c.Bind(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		_, err = domain.UserCreate(context.Background(), database, r)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}

}
