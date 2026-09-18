package main

import (
	"context"
	"fmt"

	psdPlot "github.com/ngyewch/go-dsp/psd/plot"
	"github.com/urfave/cli/v3"
)

func doGeneratePsd(ctx context.Context, cmd *cli.Command) error {
	inputFile := cmd.StringArg(inputFileArg.Name)
	outputFile := cmd.StringArg(outputFileArg.Name)
	fftSize := cmd.Int(fftSizeFlag.Name)
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

	err = psdPlot.Convert(inputFile, fftSize, step, windowFunc, outputFile, plotWidth, plotHeight)
	if err != nil {
		return err
	}

	return nil
}
