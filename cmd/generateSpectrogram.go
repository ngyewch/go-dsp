package main

import (
	"context"
	"fmt"

	spectrogramPlot "github.com/ngyewch/go-dsp/spectrogram/plot"
	"github.com/urfave/cli/v3"
)

func doGenerateSpectrogram(ctx context.Context, cmd *cli.Command) error {
	inputFile := cmd.StringArg(inputFileArg.Name)
	outputFile := cmd.StringArg(outputFileArg.Name)
	fftSize := cmd.Int(fftSizeFlag.Name)
	dbRange := cmd.Float64(dbRangeFlag.Name)
	plotWidth := cmd.Int(plotWidthFlag.Name)
	plotHeight := cmd.Int(plotHeightFlag.Name)

	windowFunc, err := getWindowFunction(ctx, cmd)
	if err != nil {
		return err
	}
	step := getStep(ctx, cmd)

	if inputFile == "" {
		return fmt.Errorf("input file not specified")
	}
	if outputFile == "" {
		outputFile = inputFile + ".png"
	}

	err = spectrogramPlot.Convert(inputFile, fftSize, step, windowFunc, outputFile, dbRange, plotWidth, plotHeight)
	if err != nil {
		return err
	}

	return nil
}
