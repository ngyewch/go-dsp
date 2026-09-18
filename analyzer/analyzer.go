package analyzer

import (
	"math"

	"github.com/madelynnblue/go-dsp/fft"
	"github.com/madelynnblue/go-dsp/window"
	"github.com/ngyewch/go-genericbuffer"
)

type Analyzer struct {
	sampleRate  float64
	nfft        int
	step        int
	windowFunc  func(int) []float64
	processFunc func(v complex128) float64
	buffer      []float64
	samples     *genericbuffer.GenericBuffer[float64]
}

func New(sampleRate float64, nfft int, step int, windowFunc func(int) []float64, processFunc func(v complex128) float64) *Analyzer {
	return &Analyzer{
		sampleRate:  sampleRate,
		nfft:        nfft,
		step:        step,
		windowFunc:  windowFunc,
		processFunc: processFunc,
		buffer:      make([]float64, nfft),
		samples:     genericbuffer.NewGenericBuffer[float64](),
	}
}

func (analyzer *Analyzer) SampleRate() float64 {
	return analyzer.sampleRate
}

func (analyzer *Analyzer) NFFT() int {
	return analyzer.nfft
}

func (analyzer *Analyzer) Step() int {
	return analyzer.step
}

func (analyzer *Analyzer) Append(samples []float64) []*AnalyzedData {
	analyzer.samples.Append(samples)
	var analyzedDataArray []*AnalyzedData
	for analyzer.samples.Len() >= analyzer.nfft {
		copy(analyzer.buffer, analyzer.samples.Peek(analyzer.nfft))
		minValue := math.NaN()
		maxValue := math.NaN()
		for _, value := range analyzer.buffer {
			if math.IsNaN(minValue) || (value < minValue) {
				minValue = value
			}
			if math.IsNaN(maxValue) || (value > maxValue) {
				maxValue = value
			}
		}
		if analyzer.windowFunc != nil {
			window.Apply(analyzer.buffer, analyzer.windowFunc)
		}
		fftResult := fft.FFTReal(analyzer.buffer)
		frequencyData := make([]float64, (analyzer.nfft/2)+1)
		for i := 0; i < (analyzer.nfft/2)+1; i++ {
			v := fftResult[i]
			frequencyData[i] = analyzer.processFunc(v)
		}
		analyzedDataArray = append(analyzedDataArray, &AnalyzedData{
			SampleRate:    analyzer.sampleRate,
			MinValue:      minValue,
			MaxValue:      maxValue,
			FrequencyData: frequencyData,
		})
		analyzer.samples.Skip(analyzer.step)
	}
	return analyzedDataArray
}
