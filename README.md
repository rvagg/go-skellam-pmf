# go-skellam-pmf

[Skellam](https://en.wikipedia.org/wiki/Skellam_distribution) PMF function in Go.

Uses a [modified Bessel function of the first kind, Iα](https://en.wikipedia.org/wiki/Bessel_function#Modified_Bessel_functions:_I%CE%B1,_K%CE%B1) to calculate the probability mass function of the Skellam distribution.

***However*** it uses a fork of components taken from https://github.com/dreading/gospecfunc which appear to be machine translations of [GNU Octave components](https://docs.octave.org/doxygen/3.8/dir_9868eb982f7926b88e9e0acf50b9194d.html) which _don't work for the cases I care about_. Within the [./amos](./amos) directory are the files I've modified to work, primarily by initialising some float64 slices that are required for some cases.
