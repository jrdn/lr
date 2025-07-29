package cmd

import (
	"fmt"
	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getParameters)
	rootCmd.AddCommand(getParameterType)
	rootCmd.AddCommand(getParameterDefault)
}

var getParameters = &cobra.Command{
	Use:   "parameters",
	Short: "Get parameter IDs",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}

		names, err := client.ParameterNames()
		if err != nil {
			return err
		}

		for _, name := range names {
			fmt.Println(name)
		}

		return nil

	},
}

var getParameterType = &cobra.Command{
	Use:   "parameter-type",
	Short: "Get parameter type",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		client, err := lr.NewClient()
		if err != nil {
			return err
		}

		ptype, err := client.ParameterType(args[0])
		if err != nil {
			return err
		}
		fmt.Println(ptype)
		return nil
	},
}

var getParameterDefault = &cobra.Command{
	Use:   "parameter-default",
	Short: "Get parameter default",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		client, err := lr.NewClient()
		if err != nil {
			return err
		}

		d, err := client.ParameterDefault(args[0])
		if err != nil {
			return err
		}
		fmt.Println(d)
		return nil
	},
}
