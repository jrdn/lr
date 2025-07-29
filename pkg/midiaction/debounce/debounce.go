package debounce

import (
	"log/slog"
	"time"
)

type DebouncedFunc func()

type DebouncerFunc func(DebouncedFunc)

const resetTime = time.Millisecond * 100
const minEvents = 5

func New(debounceTime time.Duration, name string) DebouncerFunc {
	var lastInvoke time.Time
	var count int

	return func(f DebouncedFunc) {
		if count < minEvents {
			count++
			slog.Debug("Skipping until count is met", slog.Int("count", count), slog.Int("minEvents", minEvents))
			return
		}

		if time.Since(lastInvoke) > resetTime {
			count = 0
			lastInvoke = time.Time{}
			slog.Debug("resetting")
		}

		now := time.Now()

		if now.Sub(lastInvoke) > debounceTime {
			slog.Debug("invoking")
			f()
			lastInvoke = now
		}
	}
}
