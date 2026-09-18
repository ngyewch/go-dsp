package plot

import (
	"math"

	"github.com/ngyewch/go-dsp/spectrogram"
	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
)

func ToGridXYZ(spec *spectrogram.Data) *GridXYZ {
	minValue := math.NaN()
	maxValue := math.NaN()
	result := &GridXYZ{
		SampleRate: spec.SampleRate,
		Step:       spec.Step,
	}
	for _, analyzedData := range spec.AnalyzedData {
		result.Data = append(result.Data, analyzedData.FrequencyData)
		for _, v := range analyzedData.FrequencyData {
			if math.IsNaN(minValue) || (v < minValue) {
				minValue = v
			}
			if math.IsNaN(maxValue) || (v > maxValue) {
				maxValue = v
			}
		}
	}
	result.MinValue = minValue
	result.MaxValue = maxValue
	return result
}

func ToHeatMap(gridXYZ *GridXYZ, pal palette.Palette) *plotter.HeatMap {
	if pal == nil {
		pal = defaultSpectrogramPalette
	}

	heatmap := plotter.NewHeatMap(gridXYZ, pal)
	colors := pal.Colors()
	heatmap.Underflow = colors[0]
	heatmap.Overflow = colors[len(colors)-1]
	heatmap.Rasterized = true
	return heatmap
}

func ToPlot(gridXYZ *GridXYZ, pal palette.Palette) *plot.Plot {
	heatmap := ToHeatMap(gridXYZ, pal)

	p := plot.New()
	p.X.Label.Text = "Time (seconds)"
	p.X.Tick.Marker = hplot.Ticks{
		N: 10,
	}
	p.Y.Label.Text = "Frequency (Hz)"
	p.Y.Tick.Marker = hplot.Ticks{
		N: 10,
	}
	p.Add(heatmap)

	return p
}
