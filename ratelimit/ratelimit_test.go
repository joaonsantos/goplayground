package ratelimit_test

import (
	"ratelimit"
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	rl := &ratelimit.RateLimiterImpl{
		IntervalSeconds: 1,
		MaxRequests:     2,

		Clients: ratelimit.NewClients(),
	}

	if err := rl.RateLimit("1"); err != nil {
		t.Errorf("unexpected error, got %v", err)
	}
	if err := rl.RateLimit("1"); err != nil {
		t.Errorf("unexpected error, got %v", err)
	}
	if err := rl.RateLimit("1"); err == nil {
		t.Error("expected an error, got none")
	}

	if err := rl.RateLimit("2"); err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	if err := rl.RateLimit("1"); err == nil {
		t.Error("expected an error, got none")
	}

	time.Sleep(3 * time.Second)

	if err := rl.RateLimit("1"); err != nil {
		t.Errorf("unexpected error, got %v", err)
	}

	if err := rl.RateLimit("1"); err != nil {
		t.Errorf("unexpected error, got %v", err)
	}
	if err := rl.RateLimit("1"); err == nil {
		t.Errorf("expected an error, got none")
	}
}
