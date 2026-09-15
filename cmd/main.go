package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
)

var (
	version string

	app = &cli.Command{
		Name:     "go-dsp",
		Usage:    "go-dsp",
		Version:  version,
		Commands: []*cli.Command{},
	}
)

func main() {
	err := app.Run(context.Background(), os.Args)
	if err != nil {
		slog.Error("error",
			slog.Any("err", err),
		)
		os.Exit(1)
	}
}
