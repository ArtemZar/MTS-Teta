package handlers

import (
	"fmt"
	"github.com/ArtemZar/MTS-Teta/internal/config"
	"github.com/ArtemZar/MTS-Teta/internal/middlewar"
	"github.com/ArtemZar/MTS-Teta/internal/service/profile"
	"github.com/ArtemZar/MTS-Teta/internal/tokenmanager"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	Register(router *gin.Engine)
}

type handler struct {
	Config       *config.Config
	TokenManager tokenmanager.TokenManager
}

func NewHandlers(cfg *config.Config) (*handler, error) {
	tm, err := tokenmanager.New(cfg.SignKey)
	if err != nil {
		return nil, fmt.Errorf("faild init token meneger, with error:%v", err)
	}
	return &handler{
		Config:       cfg,
		TokenManager: tm,
	}, nil
}

func (h handler) Register(router *gin.Engine) {
	auth := router.Group("/")
	auth.GET("/login", h.Login())
	auth.GET("/logout", h.Logout())

	safePoints := router.Group("/")
	safePoints.Use(middlewar.Identification(h.TokenManager))
	safePoints.GET("/i", h.I())
	safePoints.GET("/me", h.Me())
}

func (h handler) I() gin.HandlerFunc {
	return func(c *gin.Context) {
		prof := c.Request.Context().Value(middlewar.CtxKeyProfile).(profile.Profile)

		c.JSON(http.StatusOK, gin.H{
			"user_name": prof.UserName,
		})

	}
}

func (h handler) Me() gin.HandlerFunc {
	return func(c *gin.Context) {

		// ...
	}
}
