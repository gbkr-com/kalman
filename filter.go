package kalman

// Filter estimates a scalar variable from observations.
//
// See https://www.cs.unc.edu/~welch/media/pdf/kalman_intro.pdf. The naming
// of variables follows that source.
//
type Filter struct {
	X float64 // The estimated variable.
	p float64 // The variance of the estimate.
	q float64 // The process variance.
	r float64 // The measurement variance.
}

// Observe the next data.
//
func (p *Filter) Observe(z float64) {
	//
	// Revise the variance. The X value is assumed to be unchanged from the
	// last calculation.
	//
	p.p += p.q
	//
	// Incorporate the observation by calculating the Kalman gain (k) and using
	// that to adjust the estimate (X) and the variance (p).
	//
	k := p.p / (p.p + p.r)
	p.X += k * (z - p.X)
	p.p *= 1 - k
}

// New returns a Kalman filter for the given initial observation.
//
func New(x float64) *Filter {
	return &Filter{
		X: x,
		p: 1.0,
		q: 0.001,
		r: 0.001,
	}
}
