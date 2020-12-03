package oauth2

import (
	"context"
	"net/url"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"
)

func parseUrl(t *testing.T, rawUrl string) *url.URL {
	parsed, err := url.Parse(rawUrl)
	if err != nil {
		t.Error(err)
	}
	return parsed
}

func TestPrefixMatching(t *testing.T) {
	root := parseUrl(t, "https://www.example.org/foo")

	expectedGood := []*url.URL{
		parseUrl(t, "https://www.example.org/foo/bar"),
		parseUrl(t, "https://www.example.org/foo/baz"),
	}

	expectedBad := []*url.URL{
		parseUrl(t, "https://example.org/foo"),
		parseUrl(t, "https://example.org/baz"),
		parseUrl(t, "https://google.org/baz"),
	}

	limiter := &RateLimitedPrefix{prefix: root}

	for _, good := range expectedGood {
		if !limiter.Matches(good) {
			t.Errorf("%s did not match %s", root, good)
		}
	}

	for _, bad := range expectedBad {
		if limiter.Matches(bad) {
			t.Errorf("%s did not match %s", root, bad)
		}
	}
}

func TestAcquireBacklog(t *testing.T) {
	backlogFailureDetected := make(chan error)

	limiter := &RateLimitedPrefix{
		PoolC:   make(chan time.Time),
		backlog: semaphore.NewWeighted(10),
	}

	for i := 0; i <= 10; i++ {
		go func() {
			err := limiter.Acquire(context.Background())
			if err.Error() == "oauth2.RateLimitedPrefix: Request backlog full.  Try again later" {
				backlogFailureDetected <- err
			}
		}()
	}

	select {
	case <-backlogFailureDetected:
		break
	case <-time.After(1 * time.Second):
		t.Errorf("Timed out without expected error")
	}
}

func TestAcquireTimeout(t *testing.T) {
	limiter := &RateLimitedPrefix{
		PoolC:   make(chan time.Time),
		backlog: semaphore.NewWeighted(10),
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*1)
	err := limiter.Acquire(ctx)
	if err == nil {
		t.Errorf("Did not receive error")
	}
	if err != ctx.Err() {
		t.Errorf("Got unexpected error: %s", err)
	}
}
