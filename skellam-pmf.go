package skellampmf

import "math"

// SkellamPMF calculates the probability mass function (PMF) of a Skellam distribution.
//
// The Skellam distribution is the probability distribution of the difference
// of two independent Poisson random variables.
//
// Arguments:
// * k - The difference of two Poisson random variables.
// * mu1 - The expected value of the first Poisson distribution.
// * mu2 - The expected value of the second Poisson distribution.
//
// Returns:
// * A float64 representing the PMF of the Skellam distribution at k.
func SkellamPMF(k int, mu1 float64, mu2 float64) float64 {
	// Based on https://github.com/jsoares/rusty-skellam/blob/main/src/lib.rs

	// Return NaN if parameters outside range
	if math.IsNaN(mu1) || mu1 <= 0 || math.IsNaN(mu2) || mu2 <= 0 {
		return math.NaN()
	}

	// Parameterise and compute the Modified Bessel function of the first kind
	nu := float64(k)
	z := complex(2.0*math.Sqrt(mu1*mu2), 0)
	besselResult := BesselI(nu, z)

	// Compute the pmf
	return math.Exp(-(mu1 + mu2)) * math.Pow(mu1/mu2, nu/2.0) * real(besselResult)
}
