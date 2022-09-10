package handlers

import (
	"github.com/ArtemZar/MTS-Teta/internal/service/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		userName, password, ok := c.Request.BasicAuth()
		if !ok {
			c.String(http.StatusUnauthorized, "incorrect authorization header")
			return
		}

		// check credentials
		if ok := auth.CheckCredentials(h.Config, userName, password); !ok {
			c.String(http.StatusUnauthorized, "invalid credentials")
			return
		}

		tokenString, err := h.TokenManager.NewJWT(userName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
		c.SetCookie(
			"access",
			tokenString,
			int(time.Minute/1e9),
			"/",
			"localhost",
			false,
			true,
		)

		c.SetCookie(
			"refresh",
			tokenString,
			int(time.Hour/1e9),
			"/",
			"localhost",
			false,
			true,
		)

		// TODO redirect
		c.Status(http.StatusAccepted)

	}
}
