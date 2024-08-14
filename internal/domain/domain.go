package domain

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
		Use:   "domain",
		Short: "Manage DigitalOcean domains and DNS records",
	}

	cmd.AddCommand(
		listDomainsCmd(cfg),
		createDomainCmd(cfg),
		deleteDomainCmd(cfg),
		listRecordsCmd(cfg),
		createRecordCmd(cfg),
		deleteRecordCmd(cfg),
	)

	return cmd
}

func listDomainsCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all domains",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			domains, err := client.ListDomains(context.Background())
			if err != nil {
				logging.ErrorLogger.Printf("Failed to list domains: %v", err)
				return err
			}

			for _, domain := range domains {
				fmt.Printf("Name: %s, TTL: %d\n", domain.Name, domain.TTL)
			}
			return nil
		},
	}
}

func createDomainCmd(cfg *config.Config) *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			domain, err := client.CreateDomain(context.Background(), name)
			if err != nil {
				logging.ErrorLogger.Printf("Failed to create domain: %v", err)
				return err
			}

			fmt.Printf("Domain created: Name: %s, TTL: %d\n", domain.Name, domain.TTL)
			return nil
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "Domain name")
	cmd.MarkFlagRequired("name")

	return cmd
}

func deleteDomainCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [domain_name]",
		Short: "Delete a domain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			err := client.DeleteDomain(context.Background(), args[0])
			if err != nil {
				logging.ErrorLogger.Printf("Failed to delete domain: %v", err)
				return err
			}

			fmt.Printf("Domain %s deleted\n", args[0])
			return nil
		},
	}
}

func listRecordsCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "list-records [domain_name]",
		Short: "List all DNS records for a domain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			records, err := client.ListDomainRecords(context.Background(), args[0])
			if err != nil {
				logging.ErrorLogger.Printf("Failed to list domain records: %v", err)
				return err
			}

			for _, record := range records {
				fmt.Printf("ID: %d, Type: %s, Name: %s, Data: %s\n", record.ID, record.Type, record.Name, record.Data)
			}
			return nil
		},
	}
}

func createRecordCmd(cfg *config.Config) *cobra.Command {
	var recordType, name, data string
	var priority int

	cmd := &cobra.Command{
		Use:   "create-record",
		Short: "Create a new DNS record",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			record, err := client.CreateDomainRecord(context.Background(), name, recordType, data, priority)
			if err != nil {
				logging.ErrorLogger.Printf("Failed to create domain record: %v", err)
				return err
			}

			fmt.Printf("Record created: ID: %d, Type: %s, Name: %s, Data: %s\n", record.ID, record.Type, record.Name, record.Data)
			return nil
		},
	}

	cmd.Flags().StringVar(&recordType, "type", "", "Record type (e.g., A, CNAME, MX)")
	cmd.Flags().StringVar(&name, "name", "", "Record name")
	cmd.Flags().StringVar(&data, "data", "", "Record data")
	cmd.Flags().IntVar(&priority, "priority", 0, "Record priority (optional)")

	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("data")

	return cmd
}

func deleteRecordCmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-record [record_id] [domain_name]",
		Short: "Delete a DNS record",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			err := client.DeleteDomainRecord(context.Background(), args[0], args[1])
			if err != nil {
				logging.ErrorLogger.Printf("Failed to delete domain record: %v", err)
				return err
			}

			fmt.Printf("Record %s deleted from domain %s\n", args[0], args[1])
			return nil
		},
	}
}