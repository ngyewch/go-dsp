package analyzer

type AnalyzedData struct {
	SampleRate    float64   `json:"sampleRate"`
	MinValue      float64   `json:"minValue"`
	MaxValue      float64   `json:"maxValue"`
	FrequencyData []float64 `json:"frequencyData"`
}
