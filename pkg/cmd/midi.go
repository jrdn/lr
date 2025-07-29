package cmd

import (
	"github.com/jrdn/goutil/sig"
	"github.com/jrdn/lr/pkg/lr"
	"github.com/jrdn/lr/pkg/midiaction"
	"github.com/spf13/cobra"
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
	"log/slog"
)

func init() {
	rootCmd.AddCommand(midiCmd)

	midiCmd.AddCommand(midiListCmd)
	midiCmd.AddCommand(midiStartCmd)
}

var midiCmd = &cobra.Command{
	Use: "midi",
}

var midiListCmd = &cobra.Command{
	Use:   "list",
	Short: "list midi ports",
	RunE: func(cmd *cobra.Command, args []string) error {
		defer midi.CloseDriver()
		inPorts := midi.GetInPorts()
		for _, port := range inPorts {
			println(port.String())
		}
		return nil
	},
}

var midiStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start listening to midi",
	RunE: func(cmd *cobra.Command, args []string) error {
		defer midi.CloseDriver()
		inPorts := midi.GetInPorts()
		for _, port := range inPorts {
			stop, err := midi.ListenTo(port, midiCallback)
			if err != nil {
				return nil
			}
			sig.WaitUntilSignalled()
			stop()
		}
		return nil
	},
}

var client lr.Client

func midiCallback(msg midi.Message, _ int32) {
	if client == nil {
		c, err := lr.NewClient()
		if err != nil {
			panic(err)
		}
		client = c
	}

	slog.Debug("midi message", slog.String("msg", msg.String()))

	var channel, cc, val, key, velocity uint8
	switch {
	case msg.GetControlChange(&channel, &cc, &val):
		action := midiaction.GetCCAction(channel, cc)
		if action != nil {
			action(client, val)
		}
	case msg.GetNoteOn(&channel, &key, &velocity):
		action := midiaction.GetNoteAction(channel, key)
		if action != nil {
			action(client, velocity)
		}
	}
}
