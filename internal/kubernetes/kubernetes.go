package kubernetes

import (
	"context"
	"fmt"

	"github.com/felipepimentel/digitalocean-go/internal/api"
	"github.com/felipepimentel/digitalocean-go/internal/config"
	"github.com/spf13/cobra"
)

func Cmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubernetes",
		Short: "Manage DigitalOcean Kubernetes clusters",
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
		Short: "List all Kubernetes clusters",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			clusters, err := client.ListKubernetesClusters(context.Background())
			if err != nil {
				return fmt.Errorf("failed to list Kubernetes clusters: %w", err)
			}

			for _, c := range clusters {
				fmt.Printf("ID: %s, Name: %s, Region: %s, Version: %s\n", c.ID, c.Name, c.RegionSlug, c.VersionSlug)
			}
			return nil
		},
	}
}

func createCmd(cfg *config.Config) *cobra.Command {
	var name, region, version string
	var numNodes int

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new Kubernetes cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			cluster, err := client.CreateKubernetesCluster(context.Background(), name, region, version, numNodes)
			if err != nil {
				return fmt.Errorf("failed to create Kubernetes cluster: %w", err)
			}

			fmt.Printf("Kubernetes cluster created: ID: %s, Name: %s\n", cluster.ID, cluster.Name)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Name of the Kubernetes cluster")
	cmd.Flags().StringVar(&region, "region", "", "Region for the Kubernetes cluster")
	cmd.Flags().StringVar(&version, "version", "", "Kubernetes version")
	cmd.Flags().IntVar(&numNodes, "nodes", 3, "Number of nodes in the cluster")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("region")
	cmd.MarkFlagRequired("version")

	return cmd
}

func deleteCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [cluster_id]",
		Short: "Delete a Kubernetes cluster",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			err := client.DeleteKubernetesCluster(context.Background(), args[0])
			if err != nil {
				return fmt.Errorf("failed to delete Kubernetes cluster: %w", err)
			}

			fmt.Printf("Kubernetes cluster %s deleted\n", args[0])
			return nil
		},
	}
}
