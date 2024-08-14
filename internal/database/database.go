package database

import (
	"context"
	"fmt"

	"github.com/felipepimentel/digitalocean-go/internal/api"
	"github.com/felipepimentel/digitalocean-go/internal/config"
	"github.com/felipepimentel/digitalocean-go/internal/logging"
	"github.com/spf13/cobra"
)

func Cmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "Manage DigitalOcean managed databases",
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
		Short: "List all managed databases",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			databases, err := client.ListDatabases(context.Background())
			if err != nil {
				logging.ErrorLogger.Printf("Failed to list databases: %v", err)
				return err
			}

			for _, db := range databases {
				fmt.Printf("ID: %s, Name: %s, Engine: %s, Version: %s\n", db.ID, db.Name, db.EngineSlug, db.VersionSlug)
			}
			return nil
		},
	}
}

func createCmd(cfg *config.Config) *cobra.Command {
	var name, engine, version, size, region string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new managed database",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			database, err := client.CreateDatabase(context.Background(), name, engine, version, size, region)
			if err != nil {
				logging.ErrorLogger.Printf("Failed to create database: %v", err)
				return err
			}

			fmt.Printf("Database created: ID: %s, Name: %s\n", database.ID, database.Name)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Database name")
	cmd.Flags().StringVar(&engine, "engine", "", "Database engine (e.g., pg, mysql)")
	cmd.Flags().StringVar(&version, "version", "", "Database version")
	cmd.Flags().StringVar(&size, "size", "db-s-1vcpu-1gb", "Database size")
	cmd.Flags().StringVar(&region, "region", "", "Database region")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("engine")
	cmd.MarkFlagRequired("version")
	cmd.MarkFlagRequired("region")

	return cmd
}

func deleteCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [database_id]",
		Short: "Delete a managed database",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			err := client.DeleteDatabase(context.Background(), args[0])
			if err != nil {
				logging.ErrorLogger.Printf("Failed to delete database: %v", err)
				return err
			}

			fmt.Printf("Database %s deleted\n", args[0])
			return nil
		},
	}
}