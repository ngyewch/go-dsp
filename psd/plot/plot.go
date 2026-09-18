package plot

import (
	"github.com/ngyewch/go-dsp/psd"
	"go-hep.org/x/hep/hplot"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/palette/brewer"
	"gonum.org/v1/plot/plotter"
)

func ToPlot(psds []*psd.Data, pal palette.Palette, legendFunc func(i int) string) (*plot.Plot, error) {
	if pal == nil {
		pal1, err := brewer.GetPalette(brewer.TypeQualitative, "Accent", 8)
		if err != nil {
			return nil, err
		}
		pal = pal1
	}

	p := plot.New()
	p.Legend.Top = true
	p.Legend.Left = false
	p.X.Label.Text = "Frequency (Hz)"
	p.X.Tick.Marker = hplot.Ticks{
		N: 10,
	}
	p.Y.Label.Text = "dB/Hz"
	p.Y.Tick.Marker = hplot.Ticks{
		N: 10,
	}

	for i, data := range psds {
		xys := ToXYs(data)
		line, err := plotter.NewLine(xys)
		if err != nil {
			return nil, err
		}
		line.Color = pal.Colors()[i%len(pal.Colors())]
		if legendFunc != nil {
			p.Legend.Add(legendFunc(i), line)
		}
		p.Add(line)
	}

	return p, nil
}
