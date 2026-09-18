package reader

import (
	"fmt"
	"path/filepath"

	"github.com/ngyewch/go-pcm"
	"github.com/ngyewch/go-pcm/wav"
)

func PCMSourceFromFile(path string) (pcm.Source, error) {
	ext := filepath.Ext(path)
	switch ext {
	case ".wav":
		return wav.NewReader(path)
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}
}

func Float64ReaderFromFile(path string) (Float64Reader, error) {
	ext := filepath.Ext(path)
	switch ext {
	case ".wav":
		pcmSource, err := PCMSourceFromFile(path)
		if err != nil {
			return nil, err
		}
		float64Reader := NewPCMSourceFloat64Reader(pcmSource)
		return float64Reader, nil
	case ".flac":
		return FLACFloat64ReaderFromFile(path)
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}
}
