package throttler

import (
	"golang.org/x/time/rate"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// TestNewThrottler1 equality check
func TestNewThrottler1(t *testing.T) {
	testThrottler := &Throttler{transport: http.DefaultTransport,
		mapMethods: map[string]struct{}{"DELETE": {}, "GET": {}, "POST": {}, "PUT": {}},
		goodPrefix: &[]string{"/servers/.+/status"},
		badPrefix:  &[]string{"/apidomain.com/.+/routes"},
		waitLimit:  true}
	getThrottler := NewThrottler(http.DefaultTransport,
		&[]string{"GET", "POST", "PUT", "DELETE"},
		time.Minute, 0,
		&[]string{"/servers/*/status"}, &[]string{"/apidomain.com/*/routes"},
		true)
	if !reflect.DeepEqual(testThrottler, getThrottler) {
		t.Error(
			"message", "the problem is in creating a new throttler",
		)
	}
	testThrottler.goodPrefix = &[]string{}
	if reflect.DeepEqual(testThrottler, getThrottler) {
		t.Error(
			"message", "the problem is in creating a new throttler",
		)
	}
}

// TestNewThrottler2 check with nil prefixes
func TestNewThrottler2(t *testing.T) {
	testThrottler := &Throttler{transport: http.DefaultTransport,
		mapMethods: map[string]struct{}{"DELETE": {}, "GET": {}, "POST": {}, "PUT": {}},
		limiter:    rate.NewLimiter(rate.Every(time.Minute/(time.Duration(60))), 1),
		goodPrefix: new([]string),
		badPrefix:  new([]string),
		waitLimit:  true}
	getThrottler := NewThrottler(http.DefaultTransport,
		&[]string{"GET", "POST", "PUT", "DELETE"},
		time.Minute, 60,
		nil, nil,
		true)
	if !reflect.DeepEqual(testThrottler, getThrottler) {
		t.Error(
			"message", "the problem is in creating a new throttler",
		)
	}
	testThrottler.waitLimit = false
	if reflect.DeepEqual(testThrottler, getThrottler) {
		t.Error(
			"message", "the problem is in creating a new throttler",
		)
	}
}
