package main

import (
	"fmt"
	"os"

	"github.com/felipepimentel/digitalocean-go/internal/config"
	"github.com/felipepimentel/digitalocean-go/internal/droplet"
	"github.com/felipepimentel/digitalocean-go/internal/kubernetes"
	"github.com/felipepimentel/digitalocean-go/internal/vpc"
	"github.com/spf13/cobra"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "digitalocean-cli",
		Short: "A CLI for managing DigitalOcean resources",
	}

	rootCmd.AddCommand(
		droplet.Cmd(cfg),
		vpc.Cmd(cfg),
		kubernetes.Cmd(cfg),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}
