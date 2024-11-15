// Package: goTokenBucket provides a simple token bucket implementation. This can be used to rate limit requests.
package goTokenBucket

import "time"

// A TokenBucket is a simple implementation of a token bucket. It can be used to rate limit requests.
type TokenBucket struct {
	maxTokens      int
	bucket         chan bool
	refillInterval int
	refillAmount   int
}

// NewTokenBucket creates a new TokenBucket with the given parameters.
// RefillInterval is the time in milliseconds between refills.
func NewTokenBucket(maxTokens int, refillInterval int, refillAmount int, startAmount int) *TokenBucket {
	if startAmount > maxTokens {
		startAmount = maxTokens
	}
	if startAmount < 0 {
		startAmount = 0
	}
	bucket := make(chan bool, maxTokens)
	for i := 0; i < startAmount; i++ {
		bucket <- true
	}
	tokenBucket := TokenBucket{maxTokens, bucket, refillInterval, refillAmount}
	go tokenBucket.refill()
	return &tokenBucket
}

// Refill is a goroutine that refills the token bucket at the given interval.
func (t *TokenBucket) refill() {
	for {
		time.Sleep(time.Duration(t.refillInterval) * time.Millisecond)
		for i := 0; i < t.refillAmount; i++ {
			if len(t.bucket) < t.maxTokens {
				t.bucket <- true
			}
		}
	}
}

// AvailableTokens returns the number of tokens available in the bucket.
func (t *TokenBucket) AvailableTokens() int {
	return len(t.bucket)
}

// RequestTokenBlocking blocks until a token is available.
func (t *TokenBucket) RequestTokenBlocking() {
	<-t.bucket
}

// RequestTokenNonBlocking returns true if a token is available, otherwise false.
func (t *TokenBucket) RequestTokenNonBlocking() bool {
	select {
	case <-t.bucket:
		return true
	default:
		return false
	}
}
