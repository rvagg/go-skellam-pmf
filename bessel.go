package skellampmf

import (
	"math"

	"github.com/rvagg/go-skellam-pmf/amos"
)

// BesselI computes the bessel function I for complex arguments
//
// Based on github.com/dreading/gospecfunc/bessel
func BesselI(α float64, z complex128) complex128 {
	CYR := []float64{math.NaN(), 0}
	CYI := []float64{math.NaN(), 0}
	var factor complex128
	if α < 0 {
		factor = complex(2*math.Sin(math.Pi*math.Abs(α))/math.Pi, 0) * BesselK(math.Abs(α), z)
	} else {
		factor = complex(0, 0)
	}
	_, _, _, _, _, CYR, CYI, _, _ = amos.ZBESI(real(z), imag(z), math.Abs(α), 1, 1, CYR, CYI, 0, 0)
	return complex(CYR[1], CYI[1]) + factor
}

// BesselK computes the bessel function K for complex arguments
//
// Based on github.com/dreading/gospecfunc/bessel
func BesselK(α float64, z complex128) complex128 {
	CYR := []float64{math.NaN(), 0}
	CYI := []float64{math.NaN(), 0}
	_, _, _, _, _, CYR, CYI, _, _ = amos.ZBESK(real(z), imag(z), math.Abs(α), 1, 1, CYR, CYI, 0, 0)
	return complex(CYR[1], CYI[1])
}
