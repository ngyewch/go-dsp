package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/urfave/cli/v3"
)

var (
	version string

	fftSizeFlag = &cli.IntFlag{
		Name:    "fft-size",
		Usage:   "FFT size",
		Value:   1024,
		Sources: cli.EnvVars("FFT_SIZE"),
		Action: func(ctx context.Context, cmd *cli.Command, v int) error {
			if v < 2 {
				return fmt.Errorf("invalid fft-size")
			}
			if !isPowerOfTwo(v) {
				return fmt.Errorf("fft-size must be a power of two")
			}
			return nil
		},
	}
	overlapFlag = &cli.Float64Flag{
		Name:    "overlap",
		Usage:   "overlap",
		Value:   0.5,
		Sources: cli.EnvVars("OVERLAP"),
		Action: func(ctx context.Context, cmd *cli.Command, v float64) error {
			if (v < 0) || (v >= 1) {
				return fmt.Errorf("overlap must be in the range of [0, 1)")
			}
			return nil
		},
	}
	dbRangeFlag = &cli.Float64Flag{
		Name:    "db-range",
		Usage:   "db range",
		Value:   0,
		Sources: cli.EnvVars("DB_RANGE"),
		Action: func(ctx context.Context, cmd *cli.Command, v float64) error {
			if v < 0 {
				return fmt.Errorf("db-range must be 0 (unconstrained) or greater than 0")
			}
			return nil
		},
	}
	windowFunctionFlag = &cli.StringFlag{
		Name:    "window-function",
		Usage:   "window function",
		Value:   "hann",
		Sources: cli.EnvVars("WINDOW_FUNCTION"),
	}
	plotWidthFlag = &cli.IntFlag{
		Name:     "plot-width",
		Usage:    "plot width",
		Category: "Plot",
		Value:    1024,
		Sources:  cli.EnvVars("PLOT_WIDTH"),
	}
	plotHeightFlag = &cli.IntFlag{
		Name:     "plot-height",
		Usage:    "plot height",
		Category: "Plot",
		Value:    768,
		Sources:  cli.EnvVars("PLOT_HEIGHT"),
	}

	inputFileArg = &cli.StringArg{
		Name:      "input-file",
		UsageText: "(input file)",
	}
	outputFileArg = &cli.StringArg{
		Name:      "output-file",
		UsageText: "(output file)",
	}

	app = &cli.Command{
		Name:    "go-dsp",
		Usage:   "go-dsp",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "generate",
				Commands: []*cli.Command{
					{
						Name:   "spectrogram",
						Usage:  "spectrogram",
						Action: doGenerateSpectrogram,
						Arguments: []cli.Argument{
							inputFileArg,
							outputFileArg,
						},
						Flags: []cli.Flag{
							fftSizeFlag,
							overlapFlag,
							windowFunctionFlag,
							dbRangeFlag,
							plotWidthFlag,
							plotHeightFlag,
						},
					},
				},
			},
		},
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

func isPowerOfTwo(x int) bool {
	return (x > 0) && ((x & (x - 1)) == 0)
}
