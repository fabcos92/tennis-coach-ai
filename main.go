package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	config "tennis-coach-ai/cfg"
	"tennis-coach-ai/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var cfgFile string
	var cfg *config.Config

	rootCmd := &cobra.Command{
		Use:   "tennis-coach-ai",
		Short: "Tennis Coach AI API",
		Run: func(cmd *cobra.Command, args []string) {
			if err := runHTTP(cfg); err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./cfg/config.yaml", "config file (default: ../cfg/config.yaml)")

	_ = viper.BindPFlag("app.env", rootCmd.PersistentFlags().Lookup("env"))

	cfg = config.Load(cfgFile)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runHTTP(cfg *config.Config) error {
	server := http.NewServer(cfg)

	// graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("starting HTTP server on %s:%d (env=%s)\n", cfg.HTTP.Host, cfg.HTTP.Port, cfg.App.Env)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	log.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return server.Shutdown(shutdownCtx)
}
