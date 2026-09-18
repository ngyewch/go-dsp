package spectrogram

import (
	"github.com/ngyewch/go-dsp/analyzer"
)

type Generator struct {
	*analyzer.Collector
}

func NewAnalyzer(sampleRate float64, nfft int, step int, windowFunc func(int) []float64) *analyzer.Analyzer {
	return analyzer.New(sampleRate, nfft, step, windowFunc, magnitudeToDbFunc(nfft))
}

func NewGenerator(sampleRate float64, nfft int, step int, windowFunc func(int) []float64) *Generator {
	return &Generator{
		Collector: analyzer.NewCollector(NewAnalyzer(sampleRate, nfft, step, windowFunc)),
	}
}

func (generator *Generator) ToSpectrogram() *Data {
	return &Data{
		SampleRate:   generator.Analyzer().SampleRate(),
		Nfft:         generator.Analyzer().NFFT(),
		Step:         generator.Analyzer().Step(),
		AnalyzedData: generator.AnalyzedData(),
	}
}
