package cmd

import (
	"fmt"
	"github.com/0xcafed00d/joystick"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(controllerCmd)
	controllerCmd.AddCommand(controllerStartCmd)
}

var controllerCmd = &cobra.Command{
	Use: "controller",
}

var controllerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start the controller",
	RunE: func(cmd *cobra.Command, args []string) error {
		js, err := joystick.Open(8)
		if err != nil {
			return err
		}
		for {
			state, err := js.Read()
			if err != nil {
				return err
			}
			fmt.Println("buttons: %x", state.Buttons)
		}
		return nil
	},
}
