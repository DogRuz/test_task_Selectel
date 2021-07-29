package throttler

import (
	"errors"
	"golang.org/x/time/rate"
	"net/http"
)

type Throttler struct {
	transport  http.RoundTripper
	current    *http.Request
	mapMethods map[string]struct{}
	limiter    *rate.Limiter
	goodPrefix *[]string
	badPrefix  *[]string
	waitLimit  bool
}

// RoundTrip apply restrictions for url
func (t *Throttler) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	if t.limiter != nil && !checkPattern(req.URL.String(), *t.goodPrefix) {
		if _, ok := t.mapMethods[req.Method]; ok && checkPattern(req.URL.String(), *t.badPrefix) {
			if t.waitLimit {
				err := t.limiter.Wait(req.Context())
				if err != nil {
					return nil, err
				}
			} else {
				if !t.limiter.Allow() {
					return nil, errors.New("maximum requests reached")
				}
			}
		}
	}
	return t.transport.RoundTrip(req)
}
