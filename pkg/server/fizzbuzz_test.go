package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.come/gulien/fizz-buzz/pkg/fizzbuzz"
)

func TestFizzBuzzHandler(t *testing.T) {
	for i, tc := range []struct {
		int1, int2, limit, str1, str2 string
		timeout                       time.Duration
		expectStatusCode              int
		expectJSON                    string
	}{
		{
			int1:             "",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int1 required field value is empty"}`,
		},
		{
			int1:             "foo",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int1 failed to bind field value to int"}`,
		},
		{
			int1:             "1000000000000000000000000000000000000000000000000",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int1 failed to bind field value to int"}`,
		},
		{
			int1:             "1",
			int2:             "",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int2 required field value is empty"}`,
		},
		{
			int1:             "1",
			int2:             "foo",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int2 failed to bind field value to int"}`,
		},
		{
			int1:             "1",
			int2:             "1000000000000000000000000000000000000000000000000",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int2 failed to bind field value to int"}`,
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"limit required field value is empty"}`,
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "foo",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"limit failed to bind field value to int"}`,
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "1000000000000000000000000000000000000000000000000",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"limit failed to bind field value to int"}`,
		},
		{
			int1:             "0",
			int2:             "1",
			limit:            "1",
			expectStatusCode: http.StatusBadRequest,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, fizzbuzz.ErrZeroInt)
			}(),
		},
		{
			int1:             "1",
			int2:             "0",
			limit:            "1",
			expectStatusCode: http.StatusBadRequest,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, fizzbuzz.ErrZeroInt)
			}(),
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "-1",
			expectStatusCode: http.StatusBadRequest,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, fizzbuzz.ErrNegativeOrZeroLimit)
			}(),
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "0",
			expectStatusCode: http.StatusBadRequest,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, fizzbuzz.ErrNegativeOrZeroLimit)
			}(),
		},
		{
			int1:             "2",
			int2:             "3",
			limit:            "10",
			expectStatusCode: http.StatusServiceUnavailable,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, context.DeadlineExceeded)
			}(),
		},
		{
			int1:             "2",
			int2:             "3",
			limit:            "10",
			str1:             "foo",
			str2:             "bar",
			timeout:          time.Duration(10) * time.Second,
			expectStatusCode: http.StatusOK,
			expectJSON:       `["1","foo","bar","foo","5","foobar","7","foo","bar","foo"]`,
		},
	} {
		q := make(url.Values)
		q.Add("int1", tc.int1)
		q.Add("int2", tc.int2)
		q.Add("limit", tc.limit)
		q.Add("str1", tc.str1)
		q.Add("str2", tc.str2)

		e := New(tc.timeout)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/fizz-buzz?"+q.Encode(), nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)
		actualJSON := strings.TrimSpace(rec.Body.String())

		if tc.expectStatusCode != rec.Code {
			t.Errorf("test %d: expected status code %d but got %d", i, tc.expectStatusCode, rec.Code)
		}

		if !reflect.DeepEqual(tc.expectJSON, actualJSON) {
			t.Errorf("test %d: expected response '%s', but got: '%s'", i, tc.expectJSON, actualJSON)
		}
	}
}
