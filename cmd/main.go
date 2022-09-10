package main

import (
	"context"
	"fmt"
	"github.com/ArtemZar/MTS-Teta/internal/config"
	"github.com/ArtemZar/MTS-Teta/internal/httpserver"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't start the logger with error: %v", err)
	}
	defer logger.Sync() //nolint:errcheck // ok

	// init configs
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		logger.Sugar().Errorf("can't read config from file. error: %v. Will be use default configs", err)
	}
	cfg, err := config.New()
	if err != nil {
		logger.Sugar().Fatalf("can't load config with error: %v", err)
	}

	fmt.Println(cfg.Credentials)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	// running http server
	go func(ctx context.Context, cfg *config.Config, logger *zap.Logger) {
		s := httpserver.New(cfg, logger)
		if err := s.Start(); err != nil {
			logger.Sugar().Fatalf("can't start the server with error: %v", err)
		}
	}(ctx, cfg, logger)

	<-sigCh
}
