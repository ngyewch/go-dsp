package main

import (
	"context"
	"fmt"
	"io"

	"github.com/ngyewch/go-dsp/plotutils"
	"github.com/ngyewch/go-dsp/reader"
	"github.com/ngyewch/go-dsp/spectrogram"
	"github.com/ngyewch/go-dsp/spectrogram/plot"
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

	float64Reader, err := reader.Float64ReaderFromFile(inputFile)
	if err != nil {
		return err
	}
	defer func(float64Reader reader.Float64Reader) {
		_ = float64Reader.Close()
	}(float64Reader)

	generator := spectrogram.NewGenerator(float64Reader.SampleRate(), fftSize, step, windowFunc)

	for {
		channelSamples, err := float64Reader.ReadFloat64Samples(4096)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		generator.Append(channelSamples[0])
	}

	spec := generator.ToSpectrogram()
	gridXYZ := plot.ToGridXYZ(spec)
	if dbRange > 0 {
		gridXYZ.MinValue = gridXYZ.MaxValue - dbRange
	}
	p := plot.ToPlot(gridXYZ, nil)
	err = plotutils.SavePlotToFile(p, outputFile, plotWidth, plotHeight)
	if err != nil {
		return err
	}

	return nil
}
