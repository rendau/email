package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rendau/email/internal/adapters/httpapi"
	"github.com/rendau/email/internal/adapters/logger/zap"
	"github.com/rendau/email/internal/domain/core"
	"github.com/spf13/viper"
)

func Execute() {
	loadConf()

	lg, err := zap.New(viper.GetString("LOG_LEVEL"), viper.GetBool("DEBUG"), false)
	if err != nil {
		log.Fatal(err)
	}

	core := core.New(
		lg,
		viper.GetString("SMTP_ADDR"),
		viper.GetString("SMTP_AUTH_USER"),
		viper.GetString("SMTP_AUTH_PASSWORD"),
		viper.GetString("SMTP_AUTH_HOST"),
	)

	api := httpapi.New(lg, viper.GetString("HTTP_LISTEN"), core)

	lg.Infow("Starting", "http_listen", viper.GetString("HTTP_LISTEN"))

	api.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	var exitCode int

	select {
	case <-stop:
	case <-api.Wait():
		exitCode = 1
	}

	lg.Infow("Shutting down...")

	ctx, ctxCancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer ctxCancel()

	err = api.Shutdown(ctx)
	if err != nil {
		lg.Errorw("Fail to shutdown http-api", err)
		exitCode = 1
	}

	os.Exit(exitCode)
}

func loadConf() {
	viper.SetDefault("DEBUG", "false")
	viper.SetDefault("HTTP_LISTEN", ":9090")
	viper.SetDefault("LOG_LEVEL", "info")

	confFilePath := os.Getenv("CONF_PATH")
	if confFilePath == "" {
		confFilePath = "conf.yml"
	}
	viper.SetConfigFile(confFilePath)
	_ = viper.ReadInConfig()

	// env vars are in priority
	viper.AutomaticEnv()
}
