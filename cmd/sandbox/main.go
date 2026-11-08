package main

import (
	"context"
	"flag"
	"log"
	"network-sandbox/internal/config"
	"network-sandbox/internal/mininet"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configPath := flag.String(
		"config",
		"configs/default.yaml",
		"Path to sandbox configuration",
	)
	flag.Parse()

	cfg, err := config.Load(*configPath)

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Starting Sandbox: %s", cfg.Sandbox.Name)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	runner := mininet.NewRunner()

	if err := runner.Run(ctx, cfg); err != nil {
		log.Fatalf("Sandbox failed: %v", err)
	}

	log.Println("Sandbox shutdown completed.")
}
