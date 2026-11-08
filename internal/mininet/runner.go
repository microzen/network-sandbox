package mininet

import (
	"context"
	"fmt"
	"log"
	"network-sandbox/internal/config"
	"os"
	"os/exec"
)

type Runner struct {
	logger *log.Logger
}

func NewRunner() *Runner {
	return &Runner{
		logger: log.New(os.Stdout, "[MININET]", log.LstdFlags),
	}
}
func (r *Runner) Run(ctx context.Context, cfg *config.Config) error {
	if cfg.Mininet.Cleanup {
		r.logger.Println("Cleaning up previous Mininet environment...")

		cleanup := exec.CommandContext(
			ctx,
			"sudo",
			"mm",
			"-c",
		)

		cleanup.Stdout = os.Stdout
		cleanup.Stderr = os.Stderr

		if err := cleanup.Run(); err != nil {
			return fmt.Errorf("mininet cleanup failed: %w", err)
		}
	}

	args := []string{
		"mn",
		"--topo", cfg.Mininet.Topology,
		"--switch", cfg.Mininet.Switch,
		"--controller", cfg.Mininet.Controller,
	}
	if cfg.Mininet.MAC {
		args = append(args, "--mac")
	}

	r.logger.Println("Starting Mininet...")
	r.logger.Printf("Topology: %s", cfg.Mininet.Topology)

	cmd := exec.CommandContext(
		ctx,
		"sudo",
		args...,
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mininet execution failed: %w", err)
	}

	r.logger.Println("Mininet stopped.")

	return nil
}
