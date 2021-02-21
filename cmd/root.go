package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rendau/email/internal/adapters/httpapi"
	"github.com/rendau/email/internal/adapters/logger/zap"
	"github.com/rendau/email/internal/domain/core"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "email",
	Run: func(cmd *cobra.Command, args []string) {
		loadConf()

		lg, err := zap.New(viper.GetString("log_level"), viper.GetBool("debug"), false)
		if err != nil {
			log.Fatal(err)
		}

		core := core.New(
			lg,
			viper.GetString("smtp_addr"),
			viper.GetString("smtp_auth_user"),
			viper.GetString("smtp_auth_password"),
			viper.GetString("smtp_auth_host"),
		)

		api := httpapi.New(lg, viper.GetString("http_listen"), core)

		lg.Infow("Starting", "http_listen", viper.GetString("http_listen"))

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

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
}

func loadConf() {
	viper.SetDefault("debug", "false")
	viper.SetDefault("http_listen", ":9090")
	viper.SetDefault("log_level", "info")

	confFilePath := os.Getenv("CONF_PATH")
	if confFilePath == "" {
		confFilePath = "conf.yml"
	}
	viper.SetConfigFile(confFilePath)
	_ = viper.ReadInConfig()

	// env vars are in priority
	viper.AutomaticEnv()
}
