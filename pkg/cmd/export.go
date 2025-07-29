package cmd

import (
	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(openEnhanceCmd)
	rootCmd.AddCommand(openHDRMergeCmd)
	rootCmd.AddCommand(openPanoramaMergeCmd)
	rootCmd.AddCommand(openHDRPanoramaMergeCmd)
	rootCmd.AddCommand(openExportCmd)
	rootCmd.AddCommand(editInPhotoshopCmd)
	rootCmd.AddCommand(openPreferencesCmd)
}

var openEnhanceCmd = &cobra.Command{
	Use:   "enhance",
	Short: "Open the Enhance dialog for the selected photo",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.OpenEnhance()
	},
}

var openHDRMergeCmd = &cobra.Command{
	Use:   "hdr-merge",
	Short: "Open the HDR Merge dialog for the selected photos",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.OpenHDRMerge()
	},
}

var openPanoramaMergeCmd = &cobra.Command{
	Use:   "panorama-merge",
	Short: "Open the Panorama Merge dialog for the selected photos",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.OpenPanoramaMerge()
	},
}

var openHDRPanoramaMergeCmd = &cobra.Command{
	Use:   "hdr-panorama-merge",
	Short: "Open the HDR Panorama dialog for the selected photos",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.OpenHDRPanoramaMerge()
	},
}

var openExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Open the Export dialog for the selected photos",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.OpenExport()
	},
}

var editInPhotoshopCmd = &cobra.Command{
	Use:   "edit-in-photoshop",
	Short: "Open the current photo for editing in Photoshop",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.EditInPhotoshop()
	},
}

var openPreferencesCmd = &cobra.Command{
	Use:   "preferences",
	Short: "Open the preferences dialog",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.OpenPreferences()
	},
}