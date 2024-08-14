package vpc

import (
	"context"
	"fmt"

	"github.com/felipepimentel/digitalocean-go/internal/api"
	"github.com/felipepimentel/digitalocean-go/internal/config"
	"github.com/spf13/cobra"
)

func Cmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpc",
		Short: "Manage DigitalOcean VPCs",
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
		Short: "List all VPCs",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			vpcs, err := client.ListVPCs(context.Background())
			if err != nil {
				return err
			}

			for _, v := range vpcs {
				fmt.Printf("ID: %s, Name: %s, IP Range: %s\n", v.ID, v.Name, v.IPRange)
			}
			return nil
		},
	}
}

func createCmd(cfg *config.Config) *cobra.Command {
	var name, region, ipRange string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new VPC",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			vpc, err := client.CreateVPC(context.Background(), name, region, ipRange)
			if err != nil {
				return err
			}

			fmt.Printf("VPC created: ID: %s, Name: %s, IP Range: %s\n", vpc.ID, vpc.Name, vpc.IPRange)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "VPC name")
	cmd.Flags().StringVarP(&region, "region", "r", "", "VPC region")
	cmd.Flags().StringVarP(&ipRange, "ip-range", "i", "", "VPC IP range")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("region")

	return cmd
}

func deleteCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [vpc_id]",
		Short: "Delete a VPC",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			err := client.DeleteVPC(context.Background(), args[0])
			if err != nil {
				return err
			}

			fmt.Printf("VPC with ID %s deleted successfully\n", args[0])
			return nil
		},
	}
}
