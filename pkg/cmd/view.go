package cmd

import (
	"fmt"
	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(view)

}

var view = &cobra.Command{
	Use:   "view",
	Short: "view [loupe|grid|square|compare]",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("view requires a view type")
		}

		v, err := lr.ParseView(args[0])
		if err != nil {
			return err
		}

		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ShowView(v)
	},
}
