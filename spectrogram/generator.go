package spectrogram

import (
	"github.com/ngyewch/go-dsp/analyzer"
)

type Generator struct {
	sampleRate   float64
	nfft         int
	step         int
	analyzer     *analyzer.Analyzer
	analyzedData []*analyzer.AnalyzedData
}

func NewGenerator(sampleRate float64, nfft int, step int, windowFunc func(int) []float64) *Generator {
	analyzerInstance := analyzer.New(sampleRate, nfft, step, windowFunc)
	return &Generator{
		sampleRate: sampleRate,
		nfft:       nfft,
		step:       step,
		analyzer:   analyzerInstance,
	}
}

func (generator *Generator) Append(samples []float64) {
	generator.analyzedData = append(generator.analyzedData, generator.analyzer.Append(samples)...)
}

func (generator *Generator) ToSpectrogram() *Data {
	return &Data{
		SampleRate:   generator.sampleRate,
		Nfft:         generator.nfft,
		Step:         generator.step,
		AnalyzedData: generator.analyzedData,
	}
}
