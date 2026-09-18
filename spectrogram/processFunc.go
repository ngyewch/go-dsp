package spectrogram

import (
	"math"
)

func magnitudeToDbFunc(nfft int) func(complex128) float64 {
	bSi := 2 / float64(nfft)
	return func(v complex128) float64 {
		mag := (math.Sqrt((real(v)*real(v))+(imag(v)*imag(v))) * bSi) + epsilon
		return 20 * math.Log10(mag)
	}
}
