package plot

import (
	"fmt"
	"io"

	"github.com/ngyewch/go-dsp/plotutils"
	"github.com/ngyewch/go-dsp/psd"
	"github.com/ngyewch/go-dsp/reader"
)

func Convert(inputFile string, fftSize int, step int, windowFunc func(int) []float64, outputFile string, plotWidth int, plotHeight int) error {
	float64Reader, err := reader.Float64ReaderFromFile(inputFile)
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

	p, err := ToPlot(psds, nil, func(i int) string {
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
