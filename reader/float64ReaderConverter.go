package reader

import "github.com/ngyewch/go-pcm"

type Float64ReaderPCMConverter struct {
	float64Reader  Float64Reader
	encoding       pcm.Encoding
	bytesPerSample int
	bytesPerFrame  int
	putFloat64Func func(b []byte, v float64)
}

func NewFloat64ReaderPCMConverter(float64Reader Float64Reader, encoding pcm.Encoding) *Float64ReaderPCMConverter {
	return &Float64ReaderPCMConverter{
		float64Reader:  float64Reader,
		encoding:       encoding,
		bytesPerSample: encoding.BytesPerSample(),
		bytesPerFrame:  encoding.BytesPerSample() * float64Reader.NumChannels(),
		putFloat64Func: encoding.PutFloat64Func(),
	}
}

func (adapter *Float64ReaderPCMConverter) Read(b []byte) (n int, err error) {
	requestedFrameCount := len(b) / adapter.bytesPerFrame
	channelSamples, err := adapter.float64Reader.ReadFloat64Samples(requestedFrameCount)
	if err != nil {
		return 0, err
	}
	frameCount := len(channelSamples[0])
	offset := 0
	for frameNo := range frameCount {
		for channelNo := range adapter.float64Reader.NumChannels() {
			sampleBytes := b[offset : offset+adapter.bytesPerSample]
			adapter.putFloat64Func(sampleBytes, channelSamples[channelNo][frameNo])
			offset += adapter.bytesPerSample
		}
	}
	return frameCount * adapter.bytesPerFrame, nil
}
