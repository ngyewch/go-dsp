package plotutils

import "gonum.org/v1/plot/plotter"

func ZipXYs(xs []float64, ys []float64) plotter.XYs {
	arrayLen := min(len(xs), len(ys))
	points := make(plotter.XYs, arrayLen)
	for i := range arrayLen {
		points[i].X = xs[i]
		points[i].Y = ys[i]
	}
	return points
}
