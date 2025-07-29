package cmd

import (
	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rotateLeftCmd)
	rootCmd.AddCommand(rotateRightCmd)
	rootCmd.AddCommand(flipHorizontalCmd)
	rootCmd.AddCommand(flipVerticalCmd)
	rootCmd.AddCommand(rotateAspectRatioCmd)
	rootCmd.AddCommand(toggleConstrainAspectRatioCmd)
}

var rotateLeftCmd = &cobra.Command{
	Use:   "rotate-left",
	Short: "Rotate the image 90 degrees counter-clockwise",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.RotateLeft()
	},
}

var rotateRightCmd = &cobra.Command{
	Use:   "rotate-right",
	Short: "Rotate the image 90 degrees clockwise",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.RotateRight()
	},
}

var flipHorizontalCmd = &cobra.Command{
	Use:   "flip-horizontal",
	Short: "Mirror the image horizontally",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.FlipHorizontal()
	},
}

var flipVerticalCmd = &cobra.Command{
	Use:   "flip-vertical",
	Short: "Flip the image vertically",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.FlipVertical()
	},
}

var rotateAspectRatioCmd = &cobra.Command{
	Use:   "rotate-aspect-ratio",
	Short: "Rotate the aspect ratio between portrait and landscape",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.RotateAspectRatio()
	},
}

var toggleConstrainAspectRatioCmd = &cobra.Command{
	Use:   "toggle-constrain-aspect-ratio",
	Short: "Toggle aspect ratio constraint on/off",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ToggleConstrainAspectRatio()
	},
}