package handler

import (
	"blogs/server/domain"
	"blogs/server/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func signUp(database *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r request.User
		err := c.Bind(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		_, err = domain.UserCreate(c, database, r)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}

}

func login(database *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var r request.User
		err := c.Bind(&r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user, token, err := domain.Login(c, database, r)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res := gin.H{
			"user":  user,
			"token": token,
		}
		c.JSON(http.StatusOK, res)

	}
}
