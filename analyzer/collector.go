package analyzer

type Collector struct {
	analyzer     *Analyzer
	analyzedData []*AnalyzedData
}

func NewCollector(analyzer *Analyzer) *Collector {
	return &Collector{
		analyzer: analyzer,
	}
}

func (collector *Collector) Analyzer() *Analyzer {
	return collector.analyzer
}

func (collector *Collector) AnalyzedData() []*AnalyzedData {
	return collector.analyzedData
}

func (collector *Collector) Append(samples []float64) {
	collector.analyzedData = append(collector.analyzedData, collector.analyzer.Append(samples)...)
}
