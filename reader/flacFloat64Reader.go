package reader

import (
	"io"

	"github.com/mewkiz/flac"
	"github.com/ngyewch/go-genericbuffer"
)

type FLACFloat64Reader struct {
	stream      *flac.Stream
	sampleRate  float64
	numChannels int
	divisor     float64
	buffers     []*genericbuffer.GenericBuffer[float64]
	eof         bool
}

func newFLACFloat64Reader(stream *flac.Stream) *FLACFloat64Reader {
	divisor := (1 << (stream.Info.BitsPerSample - 1)) - 1
	buffers := make([]*genericbuffer.GenericBuffer[float64], stream.Info.NChannels)
	for i := range buffers {
		buffers[i] = genericbuffer.NewGenericBuffer[float64]()
	}
	return &FLACFloat64Reader{
		stream:      stream,
		sampleRate:  float64(stream.Info.SampleRate),
		numChannels: int(stream.Info.NChannels),
		divisor:     float64(divisor),
		buffers:     buffers,
	}
}

func FLACFloat64ReaderFromFile(path string) (*FLACFloat64Reader, error) {
	stream, err := flac.Open(path)
	if err != nil {
		return nil, err
	}
	return newFLACFloat64Reader(stream), nil
}

func FLACFloat64ReaderFromReader(r io.Reader) (*FLACFloat64Reader, error) {
	stream, err := flac.New(r)
	if err != nil {
		return nil, err
	}
	return newFLACFloat64Reader(stream), nil
}

func (r *FLACFloat64Reader) Close() error {
	_ = r.stream.Close()
	return nil
}

func (r *FLACFloat64Reader) SampleRate() float64 {
	return r.sampleRate
}

func (r *FLACFloat64Reader) NumChannels() int {
	return r.numChannels
}

func (r *FLACFloat64Reader) ReadFloat64Samples(samplesPerChannel int) ([][]float64, error) {
	currentBufferLen := r.buffers[0].Len()
	if currentBufferLen < samplesPerChannel {
		if !r.eof {
			for r.buffers[0].Len() < samplesPerChannel {
				frame, err := r.stream.ParseNext()
				if err != nil {
					if err == io.EOF {
						r.eof = true
						break
					}
				}
				for i := range r.numChannels {
					subframe := frame.Subframes[i]
					values := make([]float64, subframe.NSamples)
					for j, v := range subframe.Samples[:subframe.NSamples] {
						values[j] = float64(v) / r.divisor
					}
					r.buffers[i].Append(values)
				}
			}
		}
	}
	if r.buffers[0].Len() == 0 {
		return nil, io.EOF
	}
	channelSamples := make([][]float64, r.numChannels)
	for i := 0; i < r.numChannels; i++ {
		channelSamples[i] = r.buffers[i].Next(samplesPerChannel)
	}
	return channelSamples, nil
}
