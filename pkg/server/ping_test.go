package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gulien/fizz-buzz/pkg/stats"
)

func TestPingHandler(t *testing.T) {
	e := New(stats.NewInMemory(), time.Duration(1)*time.Second)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}
}
