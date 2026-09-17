package spectrogram

import (
	"github.com/ngyewch/go-dsp/analyzer"
)

type Data struct {
	SampleRate   float64
	Nfft         int
	Step         int
	AnalyzedData []*analyzer.AnalyzedData
}
