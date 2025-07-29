//go:generate go generate ./pkg/lr/...

package main

import (
	"log/slog"
	"os"

	"github.com/jrdn/lr/pkg/cmd"
)

func main() {
	level := new(slog.LevelVar)
	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
	})
	slog.SetDefault(slog.New(handler))
	level.Set(slog.LevelDebug)

	if err := cmd.RootCmd().Execute(); err != nil {
		panic(err)
	}
}
