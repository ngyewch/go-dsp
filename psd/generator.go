package psd

import (
	"math"

	"github.com/ngyewch/go-dsp/analyzer"
)

type Generator struct {
	*analyzer.Collector

	norm float64
}

func NewGenerator(sampleRate float64, nfft int, step int, windowFunc func(int) []float64) *Generator {
	analyzerInstance := analyzer.New(sampleRate, nfft, step, windowFunc, toPowerFunc())
	var norm float64
	if windowFunc != nil {
		w := windowFunc(nfft)
		for _, v := range w {
			norm += math.Pow(v, 2)
		}
	} else {
		norm = 1
	}
	return &Generator{
		Collector: analyzer.NewCollector(analyzerInstance),

		norm: norm,
	}
}

func (generator *Generator) ToPsd() *Data {
	analyzedDataArray := generator.AnalyzedData()
	binCount := len(analyzedDataArray[0].FrequencyData)
	frequencyData := make([]float64, binCount)
	for _, analyzedData := range analyzedDataArray {
		for i, v := range analyzedData.FrequencyData {
			frequencyData[i] += v
		}
	}
	for i, v := range frequencyData {
		meanValue := v / float64(len(analyzedDataArray))
		normalizedValue := meanValue / generator.norm
		frequencyData[i] = 10 * math.Log10(normalizedValue*1.5/generator.Analyzer().SampleRate())
	}
	return &Data{
		SampleRate:    generator.Analyzer().SampleRate(),
		Nfft:          generator.Analyzer().NFFT(),
		FrequencyData: frequencyData,
	}
}
