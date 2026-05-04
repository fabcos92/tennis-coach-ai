package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	config "tennis-coach-ai/cfg"
	"tennis-coach-ai/internal/application"
	"tennis-coach-ai/internal/infrastructure"
	"tennis-coach-ai/internal/infrastructure/http"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var cfgFile string

	rootCmd := &cobra.Command{
		Use:   "tennis-coach-ai",
		Short: "Tennis Coach AI API",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.Load(cfgFile)

			runHTTP(cfg)
		},
	}

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"./cfg/config.yaml",
		"config file",
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runHTTP(cfg *config.Config) {
	infra := infrastructure.New(cfg)
	app := application.NewApplication(infra)
	server := http.NewServer(cfg, app)

	// graceful shutdown signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// start server
	go func() {
		log.Printf("[SERVER] starting on %s:%d (env=%s)",
			cfg.HTTP.Host,
			cfg.HTTP.Port,
			cfg.App.Env,
		)

		if err := server.Start(); err != nil {
			log.Printf("[SERVER] http server stopped: %v", err)
		}
	}()

	// wait for shutdown signal
	<-ctx.Done()
	log.Println("[SHUTDOWN] signal received")

	// graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("[SHUTDOWN] error during shutdown: %v", err)
	}
}
