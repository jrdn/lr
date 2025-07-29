package cmd

import (
	"fmt"

	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(appVersionCmd)
}

var appVersionCmd = &cobra.Command{
	Use:   "app-version",
	Short: "Get app version string",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		ver, err := client.AppVersionString()
		if err != nil {
			return err
		}
		fmt.Println(ver)
		return nil
	},
}
