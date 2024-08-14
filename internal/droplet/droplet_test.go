package droplet

import (
	"testing"

	"github.com/felipepimentel/digitalocean-go/internal/config"
)

func TestCmd(t *testing.T) {
	cfg := &config.Config{}
	cmd := Cmd(cfg)

	if cmd.Use != "droplet" {
		t.Errorf("Expected Use to be 'droplet', got '%s'", cmd.Use)
	}

	if cmd.Short != "Manage DigitalOcean droplets" {
		t.Errorf("Expected Short to be 'Manage DigitalOcean droplets', got '%s'", cmd.Short)
	}

	if len(cmd.Commands()) != 3 {
		t.Errorf("Expected 3 subcommands, got %d", len(cmd.Commands()))
	}
}
