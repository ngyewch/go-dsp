package reader

import "io"

type Float64Reader interface {
	io.Closer

	SampleRate() float64
	NumChannels() int
	ReadFloat64Samples(samplesPerChannel int) ([][]float64, error)
}
