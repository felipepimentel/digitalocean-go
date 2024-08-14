package droplet

import (
	"context"
	"strconv"

	"github.com/felipepimentel/digitalocean-go/internal/api"
	"github.com/felipepimentel/digitalocean-go/internal/config"
	"github.com/felipepimentel/digitalocean-go/internal/logging"
	"github.com/spf13/cobra"
)

func Cmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "droplet",
		Short: "Manage DigitalOcean droplets",
	}

	cmd.AddCommand(
		listCmd(cfg),
		createCmd(cfg),
		deleteCmd(cfg),
	)

	return cmd
}

func listCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all droplets",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			droplets, err := client.ListDroplets(context.Background())
			if err != nil {
				logging.ErrorLogger.Printf("Failed to list droplets: %v", err)
				return err
			}

			for _, d := range droplets {
				logging.InfoLogger.Printf("ID: %d, Name: %s, Status: %s", d.ID, d.Name, d.Status)
			}
			return nil
		},
	}
}

func createCmd(cfg *config.Config) *cobra.Command {
	var name, region, size, image string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new droplet",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			droplet, err := client.CreateDroplet(context.Background(), name, region, size, image)
			if err != nil {
				logging.ErrorLogger.Printf("Failed to create droplet: %v", err)
				return err
			}

			logging.InfoLogger.Printf("Droplet created: ID: %d, Name: %s", droplet.ID, droplet.Name)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Droplet name")
	cmd.Flags().StringVarP(&region, "region", "r", "nyc1", "Droplet region")
	cmd.Flags().StringVarP(&size, "size", "s", "s-1vcpu-1gb", "Droplet size")
	cmd.Flags().StringVarP(&image, "image", "i", "ubuntu-20-04-x64", "Droplet image")

	cmd.MarkFlagRequired("name")

	return cmd
}

func deleteCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [droplet_id]",
		Short: "Delete a droplet",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				logging.ErrorLogger.Printf("Invalid droplet ID: %v", err)
				return err
			}

			client := api.NewClient(cfg)
			err = client.DeleteDroplet(context.Background(), id)
			if err != nil {
				logging.ErrorLogger.Printf("Failed to delete droplet: %v", err)
				return err
			}

			logging.InfoLogger.Printf("Droplet with ID %d deleted successfully", id)
			return nil
		},
	}
}
