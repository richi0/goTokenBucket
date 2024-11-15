package goTokenBucket_test

import (
	"testing"
	"time"

	"github.com/richi0/goTokenBucket"
)

func TestTokenBucket(t *testing.T) {
	tb := goTokenBucket.NewTokenBucket(10, 100, 1, 10)
	if tb.AvailableTokens() != 10 {
		t.Error("Expected 10 tokens, got", tb.AvailableTokens())
	}
	tb.RequestTokenBlocking()
	if tb.AvailableTokens() != 9 {
		t.Error("Expected 9 tokens, got", tb.AvailableTokens())
	}
	time.Sleep(110 * time.Millisecond)
	if tb.AvailableTokens() != 10 {
		t.Error("Expected 10 tokens, got", tb.AvailableTokens())
	}
	time.Sleep(110 * time.Millisecond)
	if tb.AvailableTokens() != 10 {
		t.Error("Expected 10 tokens, got", tb.AvailableTokens())
	}
}

func TestTokenBucketNonBlocking(t *testing.T) {
	tb := goTokenBucket.NewTokenBucket(10, 100, 1, 0)
	if tb.AvailableTokens() != 0 {
		t.Error("Expected 0 tokens, got", tb.AvailableTokens())
	}
	if tb.RequestTokenNonBlocking() {
		t.Error("Expected no token available")
	}
	timestampStart := time.Now().UnixMilli()
	tb.RequestTokenBlocking()
	timestampEnd := time.Now().UnixMilli()
	if timestampEnd-timestampStart < 90 {
		t.Error("Expected blocking for at least 90ms, took", timestampEnd-timestampStart, "ms")
	}
	time.Sleep(110 * time.Millisecond)
	if tb.AvailableTokens() != 1 {
		t.Error("Expected 1 token, got", tb.AvailableTokens())
	}
	if !tb.RequestTokenNonBlocking() {
		t.Error("Expected token available")
	}
}
