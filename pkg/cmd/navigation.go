package cmd

import (
	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(selectAllCmd)
	rootCmd.AddCommand(selectNoneCmd)
	rootCmd.AddCommand(nextPhotoCmd)
	rootCmd.AddCommand(prevPhotoCmd)
	rootCmd.AddCommand(showSourceCmd)
}

var selectAllCmd = &cobra.Command{
	Use:   "select-all",
	Short: "Select all photos in the grid",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.SelectAll()
	},
}

var selectNoneCmd = &cobra.Command{
	Use:   "select-none",
	Short: "Deselect all photos",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.SelectNone()
	},
}

var nextPhotoCmd = &cobra.Command{
	Use:   "next",
	Short: "Select the next photo",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.NextPhoto()
	},
}

var prevPhotoCmd = &cobra.Command{
	Use:   "prev",
	Short: "Select the previous photo",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.PrevPhoto()
	},
}

var showSourceCmd = &cobra.Command{
	Use:   "source [all_photos|trash]",
	Short: "Select the given source",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		source, err := lr.ParseSource(args[0])
		if err != nil {
			return err
		}
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ShowSource(source)
	},
}