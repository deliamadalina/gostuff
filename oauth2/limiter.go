package oauth2

import (
	"context"
	"net/url"
)

type Limiter interface {
	Matches(*url.URL) bool
	Acquire(context.Context) error
	Stop()
}
