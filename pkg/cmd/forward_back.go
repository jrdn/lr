package cmd

import (
	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(forward)
	rootCmd.AddCommand(back)
}

var forward = &cobra.Command{
	Use:   "forward",
	Short: "forward",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.GoForward()
	},
}

var back = &cobra.Command{
	Use:   "back",
	Short: "back",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.GoBack()
	},
}
