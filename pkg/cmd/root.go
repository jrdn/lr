package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lr",
	Short: "lr is an interface to Lightroom (cloud)",
}

func RootCmd() *cobra.Command {
	return rootCmd
}
