package psd

import (
	"math/cmplx"
)

func toPowerFunc() func(complex128) float64 {
	return func(v complex128) float64 {
		return real(cmplx.Conj(v)*v) * 2
	}
}
