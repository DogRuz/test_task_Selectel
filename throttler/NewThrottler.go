package throttler

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

// NewThrottler we create a new throttler
func NewThrottler(transport http.RoundTripper, methods *[]string, duration time.Duration, countRequest uint, goodPrefix *[]string, badPrefix *[]string, waitLimit bool) http.RoundTripper {
	t := Throttler{}
	t.transport = transport
	t.mapMethods = convertSliceToMap(*methods)
	if countRequest > 0 {
		rt := rate.Every(duration / (time.Duration(countRequest)))
		t.limiter = rate.NewLimiter(rt, 1)
	}
	if goodPrefix != nil {
		replace(*goodPrefix)
	} else {
		goodPrefix = new([]string)
	}
	t.goodPrefix = goodPrefix
	if badPrefix != nil {
		replace(*badPrefix)
	} else {
		badPrefix = new([]string)
	}
	t.badPrefix = badPrefix
	t.waitLimit = waitLimit
	return &t
}
