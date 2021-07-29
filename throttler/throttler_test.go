package throttler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestThrottler_RoundTripTime waitLimit true
func TestThrottler_RoundTripTime(t *testing.T) {
	getThrottler := NewThrottler(http.DefaultTransport,
		&[]string{"GET", "POST", "PUT", "DELETE"},
		time.Minute, 60,
		&[]string{"/servers/*/status"}, &[]string{"example.com/*"},
		true)
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	_, err := getThrottler.RoundTrip(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	_, err = getThrottler.RoundTrip(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestThrottler_RoundTripError waitLimit false
func TestThrottler_RoundTripError(t *testing.T) {
	getThrottler := NewThrottler(http.DefaultTransport,
		&[]string{"GET", "POST", "PUT", "DELETE"},
		time.Minute, 60,
		&[]string{"/servers/*/status"}, &[]string{"example.com/*"},
		false)
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	_, err := getThrottler.RoundTrip(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	_, err = getThrottler.RoundTrip(req)
	if err == nil {
		t.Errorf("unexpected error: max request")
	}
}
