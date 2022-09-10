package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		redirectUri := c.Query("redirect_uri")
		fmt.Println(redirectUri)
		cookies := c.Request.Cookies()
		for _, vol := range cookies {
			if vol.Name == "access" || vol.Name == "refresh" {
				vol.Value = ""
				vol.MaxAge = -1
				http.SetCookie(c.Writer, vol)
			}
		}
		c.Status(http.StatusOK)

		if redirectUri != "" {
			c.Redirect(http.StatusFound, redirectUri)
			return
		}

	}
}
