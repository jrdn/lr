package cmd

import (
	"strconv"

	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(zoomInCmd)
	rootCmd.AddCommand(zoomOutCmd)
	rootCmd.AddCommand(zoomInSomeCmd)
	rootCmd.AddCommand(zoomOutSomeCmd)
	rootCmd.AddCommand(zoom100Cmd)
	rootCmd.AddCommand(zoomFitCmd)
	rootCmd.AddCommand(zoomFillCmd)
	rootCmd.AddCommand(zoomToggleCmd)
	rootCmd.AddCommand(setFocusPointCmd)
}

var zoomInCmd = &cobra.Command{
	Use:   "zoom-in",
	Short: "Zoom in to the last level used",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ZoomIn()
	},
}

var zoomOutCmd = &cobra.Command{
	Use:   "zoom-out",
	Short: "Zoom out to the last level used",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ZoomOut()
	},
}

var zoomInSomeCmd = &cobra.Command{
	Use:   "zoom-in-some",
	Short: "Increase the zoom factor one step",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ZoomInSome()
	},
}

var zoomOutSomeCmd = &cobra.Command{
	Use:   "zoom-out-some",
	Short: "Decrease the zoom factor one step",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ZoomOutSome()
	},
}

var zoom100Cmd = &cobra.Command{
	Use:   "zoom-100",
	Short: "Zoom to 100%",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Zoom100()
	},
}

var zoomFitCmd = &cobra.Command{
	Use:   "zoom-fit",
	Short: "Set zoom mode to Fit",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ZoomFit()
	},
}

var zoomFillCmd = &cobra.Command{
	Use:   "zoom-fill",
	Short: "Set zoom mode to Fill",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ZoomFill()
	},
}

var zoomToggleCmd = &cobra.Command{
	Use:   "zoom-toggle",
	Short: "Toggle between zoomed in and zoomed out",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.ZoomToggle()
	},
}

var setFocusPointCmd = &cobra.Command{
	Use:   "focus-point <x> <y> [animate]",
	Short: "Pan the photo to the given focus point (x,y as percentages 0-1)",
	Args:  cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) error {
		x, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return err
		}
		y, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			return err
		}
		
		animate := false
		if len(args) == 3 {
			animate, err = strconv.ParseBool(args[2])
			if err != nil {
				return err
			}
		}

		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.SetFocusPoint(x, y, animate)
	},
}