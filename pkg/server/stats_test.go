package server

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gulien/fizz-buzz/pkg/stats"
)

func TestStatsHandler(t *testing.T) {
	e := New(stats.NewInMemory(), time.Duration(1)*time.Second)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	expectJSON := `{"count":0,"int1":"","int2":"","limit":"","str1":"","str2":""}`
	actualJSON := strings.TrimSpace(rec.Body.String())

	if !reflect.DeepEqual(expectJSON, actualJSON) {
		t.Errorf("expected response '%s', but got: '%s'", expectJSON, actualJSON)
	}
}
