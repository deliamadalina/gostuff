package oauth2

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"golang.org/x/sync/semaphore"
)

// RateLimitedPrefix represents a URL prefix with a particular rate limit.
type RateLimitedPrefix struct {
	Throttle *time.Ticker
	PoolC    chan time.Time
	DoneC    chan struct{}

	prefix  *url.URL
	backlog *semaphore.Weighted
}

// Matches checks if a given URL should use this limiter.
func (p *RateLimitedPrefix) Matches(src *url.URL) bool {
	return strings.HasPrefix(src.String(), p.prefix.String())
}

// Acquire attempts to obtain a token from the rate limited bucket for making a
// request against this prefix.
func (p *RateLimitedPrefix) Acquire(ctx context.Context) error {
	if ok := p.backlog.TryAcquire(defaultWeight); !ok {
		return errors.New("oauth2.RateLimitedPrefix: Request backlog full.  Try again later")
	}
	defer p.backlog.Release(defaultWeight)
	select {
	case <-p.PoolC:
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

// Stop shuts down the accumulator preventing further requests.
func (p *RateLimitedPrefix) Stop() {
	p.Throttle.Stop()
	close(p.DoneC)
}

// NewRateLimitedPrefix creates a rate limiter for a given URL prefix.
func NewRateLimitedPrefix(requestLimit, poolSize int, backlogSize int64, prefix *url.URL) *RateLimitedPrefix {
	done := make(chan struct{})
	pool := make(chan time.Time, poolSize)
	throttle := time.NewTicker(time.Second / time.Duration(requestLimit))
	go func() {
		for t := range throttle.C {
			select {
			case pool <- t:
			case <-done:
				return
			}
		}
	}()

	return &RateLimitedPrefix{
		Throttle: throttle,
		PoolC:    pool,
		DoneC:    done,
		prefix:   prefix,
		backlog:  semaphore.NewWeighted(backlogSize),
	}
}
