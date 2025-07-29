package midiaction

import (
	"github.com/go-vgo/robotgo"
	"github.com/jrdn/lr/pkg/lr"
	"github.com/jrdn/lr/pkg/midiaction/debounce"
	"log/slog"
	"time"
)

// TODO this debouncer sucks because it only invokes the function when it stops getting called
// really I want to invoke it when it's first called, and periodically at some maximum rate after that

var (
	exposureDebounce   = debounce.New(50*time.Millisecond, "exposure")
	contrastDebounce   = debounce.New(25*time.Millisecond, "contrast")
	highlightsDebounce = debounce.New(25*time.Millisecond, "highlights")
	shadowsDebounce    = debounce.New(25*time.Millisecond, "shadows")
	whitesDebounce     = debounce.New(25*time.Millisecond, "whites")
	blacksDebounce     = debounce.New(25*time.Millisecond, "blacks")
)

const (
	exposureKnob     = 50
	exposureFineKnob = 54
	contrastKnob     = 58
	highlightsKnob   = 51
	shadowsKnob      = 55
	whitesKnob       = 59
	blacksKnob       = 63
)

var (
	tempDebounce       = debounce.New(50*time.Millisecond, "temp")
	tintDebounce       = debounce.New(50*time.Millisecond, "tint")
	vibranceDebounce   = debounce.New(50*time.Millisecond, "vibrance")
	saturationDebounce = debounce.New(50*time.Millisecond, "saturation")
)

const (
	tempKnob       = 52
	tintKnob       = 56
	vibranceKnob   = 60
	saturationKnob = 64
)

var (
	photoDebounce = debounce.New(200*time.Millisecond, "photo")
	zoomDebounce  = debounce.New(50*time.Millisecond, "zoom")
)

const (
	photoKnob = 62
	zoomKnob  = 65
)

var panDebounce = debounce.New(50*time.Millisecond, "pan")

const (
	panXKnob = 47
	panYKnob = 48
)

var (
	focusX float64 = 0.5
	focusY float64 = 0.5
)

func logErr(err error) {
	if err != nil {
		slog.Error("lr error", err)
	}
}

var upDownDebounce = debounce.New(50*time.Millisecond, "updown")

var ccActionMap CCActionMapping = map[uint8]map[uint8]MidiAction{
	0: {
		38: func(c lr.Client, value uint8) {
			upDownDebounce(func() {
				if value > 64 {
					logErr(robotgo.KeyTap("up"))
				} else {
					logErr(robotgo.KeyTap("down"))
				}
			})
		},
		57: func(c lr.Client, value uint8) {
			upDownDebounce(func() {
				if value > 64 {
					logErr(robotgo.KeyTap("up"))
				} else {
					logErr(robotgo.KeyTap("down"))
				}
			})
		},

		exposureKnob: func(c lr.Client, value uint8) {
			exposureDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Exposure2012", 0.1))
				} else {
					logErr(c.Decrement("Exposure2012", 0.1))
				}
			})
		},
		exposureFineKnob: func(c lr.Client, value uint8) {
			exposureDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Exposure2012", 0.01))
				} else {
					logErr(c.Decrement("Exposure2012", 0.01))
				}
			})
		},

		contrastKnob: func(c lr.Client, value uint8) {
			contrastDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Contrast2012", 1))
				} else {
					logErr(c.Decrement("Contrast2012", 1))
				}
			})
		},

		highlightsKnob: func(c lr.Client, value uint8) {
			highlightsDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Highlights2012", 1))
				} else {
					logErr(c.Decrement("Highlights2012", 1))
				}
			})
		},
		shadowsKnob: func(c lr.Client, value uint8) {
			shadowsDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Shadows2012", 1))
				} else {
					logErr(c.Decrement("Shadows2012", 1))
				}
			})
		},
		whitesKnob: func(c lr.Client, value uint8) {
			whitesDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Whites2012", 1))
				} else {
					logErr(c.Decrement("Whites2012", 1))
				}
			})
		},
		blacksKnob: func(c lr.Client, value uint8) {
			blacksDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Blacks2012", 1))
				} else {
					logErr(c.Decrement("Blacks2012", 1))
				}
			})
		},

		tempKnob: func(c lr.Client, value uint8) {
			tempDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Temperature", 10))
				} else {
					logErr(c.Decrement("Temperature", 10))
				}
			})
		},
		tintKnob: func(c lr.Client, value uint8) {
			tintDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Tint", 1))
				} else {
					logErr(c.Decrement("Tint", 1))
				}
			})
		},
		vibranceKnob: func(c lr.Client, value uint8) {
			vibranceDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Vibrance", 1))
				} else {
					logErr(c.Decrement("Vibrance", 1))
				}
			})
		},
		saturationKnob: func(c lr.Client, value uint8) {
			saturationDebounce(func() {
				if value > 64 {
					logErr(c.Increment("Saturation", 1))
				} else {
					logErr(c.Decrement("Saturation", 1))
				}
			})
		},

		photoKnob: func(c lr.Client, value uint8) {
			photoDebounce(func() {
				if value > 64 {
					logErr(c.NextPhoto())
				} else {
					logErr(c.PrevPhoto())
				}
			})
		},

		zoomKnob: func(c lr.Client, value uint8) {
			zoomDebounce(func() {
				if value > 64 {
					logErr(c.ZoomInSome())
				} else {
					logErr(c.ZoomOutSome())
				}
			})
		},

		panXKnob: func(c lr.Client, value uint8) {
			panDebounce(func() {
				focusX = float64(value) / 127
				logErr(c.SetFocusPoint(focusX, focusY, true))
			})
		},
		panYKnob: func(c lr.Client, value uint8) {
			panDebounce(func() {
				focusY = float64(value) / 127
				logErr(c.SetFocusPoint(focusX, focusY, true))
			})
		},
	},
}

var view = lr.ViewGrid

var noteActionMap NoteActionMapping = map[uint8]map[uint8]MidiAction{
	1: {
		exposureKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Exposure2012"))
		},
		contrastKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Contrast2012"))
		},
		highlightsKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Highlights2012"))
		},
		shadowsKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Shadows2012"))
		},
		whitesKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Whites2012"))
		},
		blacksKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Blacks2012"))
		},

		tempKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Temperature"))
		},
		tintKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Tint"))
		},
		vibranceKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Vibrance"))
		},
		saturationKnob: func(c lr.Client, value uint8) {
			logErr(c.ResetToDefault("Saturation"))
		},

		zoomKnob: func(c lr.Client, value uint8) {
			logErr(c.ZoomToggle())
		},

		53: func(c lr.Client, value uint8) {
			logErr(c.FlagPick())
		},
		57: func(c lr.Client, value uint8) {
			logErr(c.FlagUnflag())
		},
		61: func(c lr.Client, value uint8) {
			logErr(c.FlagReject())
		},

		36: func(c lr.Client, value uint8) {
			logErr(c.RatingIncrease())
		},
		40: func(c lr.Client, value uint8) {
			logErr(c.RatingDecrease())
		},

		photoKnob: func(c lr.Client, value uint8) {
			logErr(c.ShowView(view))

			// toggle the view
			if view == lr.ViewGrid {
				view = lr.ViewLoupe
			} else {
				view = lr.ViewGrid
			}
		},

		43: func(c lr.Client, value uint8) {
			logErr(c.Undo())
		},
		44: func(c lr.Client, value uint8) {
			logErr(c.Redo())
		},
	},

	2: {
		31: func(c lr.Client, value uint8) {
			logErr(c.SetAutoTone())
		},
	},

	3: {
		31: func(c lr.Client, value uint8) {
			logErr(c.SetAutoTone())
		},
	},
}

func GetNoteAction(channel, note uint8) MidiAction {
	if channelMap, ok := noteActionMap[channel]; ok {
		if action, ok := channelMap[note]; ok {
			return action
		}
	}
	return nil
}

func GetCCAction(channel, cc uint8) MidiAction {
	if channelMap, ok := ccActionMap[channel]; ok {
		if action, ok := channelMap[cc]; ok {
			return action
		}
	}
	return nil
}

type CCActionMapping map[uint8]map[uint8]MidiAction
type NoteActionMapping map[uint8]map[uint8]MidiAction
type MidiAction func(c lr.Client, value uint8)
