package oauth2

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/internal"
)

const defaultWeight int64 = 1

// oauth2.NewClient will use the http.Client associated with the context to
// pull its underlying Transport from.  We use this to layer another rate
// limited transport into the chain.  The resulting chain ends up being [OAuth2
// Transport] -> [Rate Limited Transport] -> [http.RoundTripper].
func WithRateLimitTransport(ctx context.Context, limiters []Limiter, base http.RoundTripper) context.Context {
	return WithTransport(ctx, NewRateLimitedTransport(limiters, base))
}

// oauth2.NewClient will use the http.Client associated with the context to
// pull its underlying Transport from.  We use this to layer another transport
// into the chain.  The resulting chain ends up being [OAuth2 Transport] ->
// [Our Transport] -> [http.RoundTripper].
func WithTransport(ctx context.Context, transport http.RoundTripper) context.Context {
	return context.WithValue(ctx, internal.HTTPClient, &http.Client{
		Transport: transport,
	})
}

// RateLimitedTransport wraps the http.RoundTripper transport interface and can
// enforce a rate limit on the number of HTTP requests that can be made.  The
// rate limits are defined on a per URL prefix basis.
type RateLimitedTransport struct {
	Limiters []Limiter
	Base     http.RoundTripper
}

// NewRateLimitedTransport creates the rate limited transport.  The rate
// limited transport takes a number of Limiter objects which impose
// their own rate limits for a URL prefix.  The transport will attempt to
// acquire a request token from the first limiter that matches the URL being
// requested.  If no limiters match, no rate limiting is imposed.
func NewRateLimitedTransport(limiters []Limiter, base http.RoundTripper) *RateLimitedTransport {
	return &RateLimitedTransport{
		Base:     base,
		Limiters: limiters,
	}
}

// Stop the rate limiters associated with this transport.   No further requests
// can be made.
func (t *RateLimitedTransport) Stop() {
	for _, l := range t.Limiters {
		l.Stop()
	}
}

// RoundTrip acquires request from the rate limited pool and makes the request
// using the underlying http.RoundTrip interface.
func (t *RateLimitedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for _, l := range t.Limiters {
		if l.Matches(req.URL) {
			if err := l.Acquire(req.Context()); err != nil {
				return nil, err
			}
			break
		}
	}
	return t.base().RoundTrip(req)
}

// CancelRequest allows for a request made to the base client to be cancelled.
func (t *RateLimitedTransport) CancelRequest(req *http.Request) {
	type canceler interface {
		CancelRequest(*http.Request)
	}
	if cr, ok := t.base().(canceler); ok {
		cr.CancelRequest(req)
	}
}

func (t *RateLimitedTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}
