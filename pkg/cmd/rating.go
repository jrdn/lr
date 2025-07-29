package cmd

import (
	"github.com/jrdn/lr/pkg/lr"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(flagPickCmd)
	rootCmd.AddCommand(flagRejectCmd)
	rootCmd.AddCommand(flagUnflagCmd)
	rootCmd.AddCommand(flagIncreaseCmd)
	rootCmd.AddCommand(flagDecreaseCmd)
	rootCmd.AddCommand(ratingIncreaseCmd)
	rootCmd.AddCommand(ratingDecreaseCmd)
	rootCmd.AddCommand(ratingClearCmd)
	rootCmd.AddCommand(rating0Cmd)
	rootCmd.AddCommand(rating1Cmd)
	rootCmd.AddCommand(rating2Cmd)
	rootCmd.AddCommand(rating3Cmd)
	rootCmd.AddCommand(rating4Cmd)
	rootCmd.AddCommand(rating5Cmd)
}

var flagPickCmd = &cobra.Command{
	Use:   "flag-pick",
	Short: "Set the flag to 'picked'",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.FlagPick()
	},
}

var flagRejectCmd = &cobra.Command{
	Use:   "flag-reject",
	Short: "Set the flag to 'rejected'",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.FlagReject()
	},
}

var flagUnflagCmd = &cobra.Command{
	Use:   "flag-unflag",
	Short: "Set the flag to 'unflagged'",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.FlagUnflag()
	},
}

var flagIncreaseCmd = &cobra.Command{
	Use:   "flag-increase",
	Short: "Increase the flag ('unflagged' to 'picked' or 'rejected' to 'unflagged')",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.FlagIncrease()
	},
}

var flagDecreaseCmd = &cobra.Command{
	Use:   "flag-decrease",
	Short: "Decrease the flag ('picked' to 'unflagged' or 'unflagged' to 'rejected')",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.FlagDecrease()
	},
}

var ratingIncreaseCmd = &cobra.Command{
	Use:   "rating-increase",
	Short: "Increase the rating",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.RatingIncrease()
	},
}

var ratingDecreaseCmd = &cobra.Command{
	Use:   "rating-decrease",
	Short: "Decrease the rating",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.RatingDecrease()
	},
}

var ratingClearCmd = &cobra.Command{
	Use:   "rating-clear",
	Short: "Clear the rating (set to 0)",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.RatingClear()
	},
}

var rating0Cmd = &cobra.Command{
	Use:   "rating-0",
	Short: "Set the rating to 0",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Rating0()
	},
}

var rating1Cmd = &cobra.Command{
	Use:   "rating-1",
	Short: "Set the rating to 1",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Rating1()
	},
}

var rating2Cmd = &cobra.Command{
	Use:   "rating-2",
	Short: "Set the rating to 2",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Rating2()
	},
}

var rating3Cmd = &cobra.Command{
	Use:   "rating-3",
	Short: "Set the rating to 3",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Rating3()
	},
}

var rating4Cmd = &cobra.Command{
	Use:   "rating-4",
	Short: "Set the rating to 4",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Rating4()
	},
}

var rating5Cmd = &cobra.Command{
	Use:   "rating-5",
	Short: "Set the rating to 5",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := lr.NewClient()
		if err != nil {
			return err
		}
		return client.Rating5()
	},
}