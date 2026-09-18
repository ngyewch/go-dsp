package plot

import (
	"github.com/ngyewch/go-dsp/plotutils"
	"github.com/ngyewch/go-dsp/psd"
	"gonum.org/v1/plot/plotter"
)

func ToXYs(data *psd.Data) plotter.XYs {
	frequencies := data.Frequencies()
	return plotutils.ZipXYs(frequencies, data.FrequencyData)
}
