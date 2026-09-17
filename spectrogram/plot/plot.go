package plot

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/ngyewch/go-dsp/spectrogram"
	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
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

func SavePlotToFile(path string, p *plot.Plot, width int, height int) error {
	ext := filepath.Ext(path)
	switch ext {
	case ".png":
	default:
		return fmt.Errorf("unsupported file type: %s", ext)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	img := vgimg.New(vg.Length(width)*vg.Inch/vg.Length(vgimg.DefaultDPI), vg.Length(height)*vg.Inch/vg.Length(vgimg.DefaultDPI))
	dc := draw.New(img)
	p.Draw(dc)

	switch ext {
	case ".png":
		png := vgimg.PngCanvas{
			Canvas: img,
		}
		_, err = png.WriteTo(f)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unsupported file type: %s", ext)
	}
}
