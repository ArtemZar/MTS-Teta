package httpserver

import (
	"github.com/ArtemZar/MTS-Teta/internal/config"
	"github.com/ArtemZar/MTS-Teta/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpServer struct {
	Config *config.Config
	Logger *zap.Logger
	Router *gin.Engine
}

func New(config *config.Config, logger *zap.Logger) *HttpServer {
	return &HttpServer{
		Config: config,
		Logger: logger,
		Router: gin.Default(),
	}
}

func (s *HttpServer) Start() error {
	handler, err := handlers.NewHandlers(s.Config)
	if err != nil {
		s.Logger.Sugar().Errorf("fail init handlers on server startup, error: %v", err)
		return err
	}
	handler.Register(s.Router)
	return s.Router.Run(s.Config.SrvConfig.Addr)
}
