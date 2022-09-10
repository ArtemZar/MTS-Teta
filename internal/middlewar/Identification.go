package middlewar

import (
	"context"
	"github.com/ArtemZar/MTS-Teta/internal/service/profile"
	"github.com/ArtemZar/MTS-Teta/internal/tokenmanager"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const CtxKeyProfile ctxKey = iota

type ctxKey int8

func Identification(tm tokenmanager.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var NeedUpdateCookies bool
		u := profile.Profile{}

		// check for jwt cookies
		var jwtCookie *http.Cookie
		jwtCookie, err := c.Request.Cookie("access")
		if err != nil {
			jwtCookie, err = c.Request.Cookie("refresh")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			NeedUpdateCookies = true
		}

		userName, _ := tm.Parse(jwtCookie.Value)
		u.UserName = userName

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), CtxKeyProfile, u))

		//  update JWT cookie
		if NeedUpdateCookies {
			tokenString, err := tm.NewJWT(userName)
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

			refreshCookie, err := c.Request.Cookie("refresh")
			if err != nil {
				c.SetCookie(
					"refresh",
					tokenString,
					int(time.Hour/1e9),
					"/",
					"localhost",
					false,
					true,
				)
			}
			refreshCookie.Value = tokenString
			refreshCookie.MaxAge = refreshCookie.MaxAge + int(time.Minute/1e9)
		}

	}

}
