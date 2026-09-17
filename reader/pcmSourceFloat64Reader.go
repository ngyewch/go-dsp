package reader

import (
	"fmt"
	"io"

	"github.com/ngyewch/go-pcm"
)

type PCMSourceFloat64Reader struct {
	source pcm.Source
}

func NewPCMSourceFloat64Reader(source pcm.Source) *PCMSourceFloat64Reader {
	return &PCMSourceFloat64Reader{
		source: source,
	}
}

func (r *PCMSourceFloat64Reader) Close() error {
	closer, ok := r.source.(io.Closer)
	if ok && (closer != nil) {
		_ = closer.Close()
	}
	return nil
}

func (r *PCMSourceFloat64Reader) SampleRate() float64 {
	return float64(r.source.SampleRate())
}

func (r *PCMSourceFloat64Reader) NumChannels() int {
	return int(r.source.NumChannels())
}

func (r *PCMSourceFloat64Reader) ReadFloat64Samples(samplesPerChannel int) ([][]float64, error) {
	numChannels := r.source.NumChannels()
	sourceEncoding := r.source.Encoding()
	toFloat64 := sourceEncoding.Float64Func()
	bytesPerSample := sourceEncoding.BytesPerSample()
	bytesPerFrame := int(numChannels) * bytesPerSample
	buffer := make([]byte, bytesPerFrame*samplesPerChannel)
	readLen, err := r.source.Read(buffer)
	if err != nil {
		return nil, err
	}
	actualSamplesPerChannel := readLen / bytesPerFrame
	channelSamples := make([][]float64, numChannels)
	for channelNo := range numChannels {
		channelSamples[channelNo] = make([]float64, actualSamplesPerChannel)
	}
	offset := 0
	for sampleNo := range actualSamplesPerChannel {
		for channelNo := range numChannels {
			channelSamples[channelNo][sampleNo] = toFloat64(buffer[offset : offset+bytesPerSample])
			offset += bytesPerSample
		}
	}
	if (readLen % bytesPerFrame) != 0 {
		return channelSamples, fmt.Errorf("incomplete frame read")
	}
	return channelSamples, nil
}
