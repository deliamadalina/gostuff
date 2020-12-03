package main

import "net/http"

type newRoundTripper struct {
	r http.RoundTripper
}

func (rt newRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("X-Auth-Key", readEnv().api_key)
	r.Header.Add("X-Auth-Email", readEnv().email)
	return rt.r.RoundTrip(r)
}
