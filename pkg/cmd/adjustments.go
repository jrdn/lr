package cmd

import (
	"fmt"
	"strconv"

	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setAutoToneCmd)
	rootCmd.AddCommand(resetToDefaultCmd)
	rootCmd.AddCommand(incrementCmd)
	rootCmd.AddCommand(decrementCmd)
	rootCmd.AddCommand(getValueCmd)
	rootCmd.AddCommand(setValueCmd)
	rootCmd.AddCommand(undoCmd)
	rootCmd.AddCommand(redoCmd)
}

var setAutoToneCmd = &cobra.Command{
	Use:   "auto-tone",
	Short: "Apply Auto Tone to the current photo",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.SetAutoTone()
	},
}

var resetToDefaultCmd = &cobra.Command{
	Use:   "reset-to-default <parameter>",
	Short: "Reset a single adjustment to its default value",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ResetToDefault(args[0])
	},
}

var incrementCmd = &cobra.Command{
	Use:   "increment <parameter> [amount]",
	Short: "Increment the value of an adjustment",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var amount float64 = 0 // Default amount
		if len(args) == 2 {
			var err error
			amount, err = strconv.ParseFloat(args[1], 64)
			if err != nil {
				return err
			}
		}

		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Increment(args[0], amount)
	},
}

var decrementCmd = &cobra.Command{
	Use:   "decrement <parameter> [amount]",
	Short: "Decrement the value of an adjustment",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var amount float64 = 0 // Default amount
		if len(args) == 2 {
			var err error
			amount, err = strconv.ParseFloat(args[1], 64)
			if err != nil {
				return err
			}
		}

		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Decrement(args[0], amount)
	},
}

var getValueCmd = &cobra.Command{
	Use:   "get-value <parameter>",
	Short: "Get the current value of an adjustment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		value, err := client.GetValue(args[0])
		if err != nil {
			return err
		}
		fmt.Println(value)
		return nil
	},
}

var setValueCmd = &cobra.Command{
	Use:   "set-value <parameter> <value>",
	Short: "Set the value of an adjustment",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Parse the value - could be string, number, or boolean
		var paramValue *lr.ParameterValue
		
		// Try to parse as float first
		if floatVal, err := strconv.ParseFloat(args[1], 64); err == nil {
			paramValue = &lr.ParameterValue{Number: &floatVal}
		} else if boolVal, err := strconv.ParseBool(args[1]); err == nil {
			paramValue = &lr.ParameterValue{Boolean: &boolVal}
		} else {
			// Treat as string
			paramValue = &lr.ParameterValue{String: &args[1]}
		}

		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.SetValue(args[0], paramValue)
	},
}

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last operation",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Undo()
	},
}

var redoCmd = &cobra.Command{
	Use:   "redo",
	Short: "Redo the last undone operation",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Redo()
	},
}