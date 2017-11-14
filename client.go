package wego

import (
	"time"

	"github.com/cenkalti/backoff"
)

// NewBackOff declare retry exponential back off with max interval time and max elapsed time in second
func NewBackOff(maxInterval, maxElapsedTime int) *backoff.ExponentialBackOff {
	expBackOff := backoff.NewExponentialBackOff()
	expBackOff.MaxElapsedTime = time.Duration(maxElapsedTime) * time.Second
	expBackOff.MaxInterval = time.Duration(maxInterval) * time.Second
	return expBackOff
}