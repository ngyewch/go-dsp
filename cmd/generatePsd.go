package main

import (
	"context"
	"fmt"
	"io"

	"github.com/ngyewch/go-dsp/plotutils"
	"github.com/ngyewch/go-dsp/psd"
	psdPlot "github.com/ngyewch/go-dsp/psd/plot"
	"github.com/ngyewch/go-dsp/reader"
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

	float64Reader, err := newFloat64Reader(inputFile)
	if err != nil {
		return err
	}
	defer func(float64Reader reader.Float64Reader) {
		_ = float64Reader.Close()
	}(float64Reader)

	generators := make([]*psd.Generator, float64Reader.NumChannels())
	for i := range float64Reader.NumChannels() {
		generators[i] = psd.NewGenerator(float64Reader.SampleRate(), fftSize, step, windowFunc)
	}

	for {
		channelSamples, err := float64Reader.ReadFloat64Samples(4096)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		for i := range float64Reader.NumChannels() {
			generators[i].Append(channelSamples[i])
		}
	}

	psds := make([]*psd.Data, float64Reader.NumChannels())
	for i := range float64Reader.NumChannels() {
		psds[i] = generators[i].ToPsd()
	}

	p, err := psdPlot.ToPlot(psds, nil, func(i int) string {
		return fmt.Sprintf("Channel %d", i)
	})
	if err != nil {
		return err
	}

	err = plotutils.SavePlotToFile(p, outputFile, plotWidth, plotHeight)
	if err != nil {
		return err
	}

	return nil
}
