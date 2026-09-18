package spectrogram

import "math"

var (
	epsilon = math.Nextafter(1.0, 2.0) - 1.0
)
