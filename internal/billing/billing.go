package billing

import (
	"context"
	"fmt"

	"github.com/felipepimentel/digitalocean-go/internal/api"
	"github.com/felipepimentel/digitalocean-go/internal/config"
	"github.com/felipepimentel/digitalocean-go/internal/logging"
	"github.com/spf13/cobra"
)

func Cmd(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "billing",
		Short: "Show billing information",
		RunE: func(cmd *cobra.Command, args []string) error {
			client := api.NewClient(cfg)
			billing, err := client.GetBillingInfo(context.Background())
			if err != nil {
				logging.ErrorLogger.Printf("Failed to get billing information: %v", err)
				return err
			}

			fmt.Printf("Month-to-date usage: $%.2f\n", billing.MonthToDateUsage)
			fmt.Printf("Account balance: $%.2f\n", billing.AccountBalance)
			fmt.Printf("Month-to-date balance: $%.2f\n", billing.MonthToDateBalance)

			return nil
		},
	}
}