package plot

type GridXYZ struct {
	SampleRate float64
	Step       int
	Data       [][]float64
	MinValue   float64
	MaxValue   float64
}

func (gridXYZ GridXYZ) Min() float64 {
	return gridXYZ.MinValue
}

func (gridXYZ GridXYZ) Max() float64 {
	return gridXYZ.MaxValue
}

func (gridXYZ GridXYZ) Dims() (int, int) {
	return len(gridXYZ.Data), len(gridXYZ.Data[0])
}

func (gridXYZ GridXYZ) Z(c int, r int) float64 {
	return gridXYZ.Data[c][r]
}

func (gridXYZ GridXYZ) X(c int) float64 {
	return (float64(c) * float64(gridXYZ.Step)) / gridXYZ.SampleRate
}

func (gridXYZ GridXYZ) Y(r int) float64 {
	v := float64(r) / float64(len(gridXYZ.Data[0]))
	return v * gridXYZ.SampleRate / 2
}
