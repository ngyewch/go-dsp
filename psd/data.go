package psd

type Data struct {
	SampleRate    float64
	Nfft          int
	FrequencyData []float64
}

func (data *Data) Frequencies() []float64 {
	binCount := len(data.FrequencyData)
	frequencies := make([]float64, binCount)
	frequencyStep := (data.SampleRate / 2) / float64(binCount-1)
	for i := range frequencies {
		frequencies[i] = float64(i) * frequencyStep
	}
	return frequencies
}
