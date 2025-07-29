package cmd

import (
	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clearFiltersCmd)
	rootCmd.AddCommand(showFiltersCmd)
}

var clearFiltersCmd = &cobra.Command{
	Use:   "clear-filters",
	Short: "Clear the search filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ClearFilters()
	},
}

var showFiltersCmd = &cobra.Command{
	Use:   "show-filters",
	Short: "Show/hide the search filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ShowFilters()
	},
}