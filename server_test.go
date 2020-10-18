package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	api "github.com/charliebillen/time-api"
)

func TestServerRouting(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		status int
	}{
		{name: "It responds OK for a handled path", path: "/time", status: http.StatusOK},
		{name: "It responds Not Found for an uhandled path", path: "/404", status: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr, rq := setupRequest(tt.path)

			srv := &api.Server{api.DefaultTimeProvider}
			srv.ServeHTTP(rr, rq)

			assertHasStatus(t, rr, tt.status)
		})
	}
}

func TestServerResponse(t *testing.T) {
	rr, rq := setupRequest("/time")

	timeProvider := func() time.Time {
		return setupTime(12, 23, 55)
	}

	srv := &api.Server{timeProvider}
	srv.ServeHTTP(rr, rq)

	want := buildResponseFrom(strings.NewReader(`{"hour":12, "minute":23, "second":55}`))
	got := buildResponseFrom(rr.Body)

	assertDeepEqual(t, want, got)
	assertHasContentType(t, rr, "application/json")
}

func setupRequest(path string) (*httptest.ResponseRecorder, *http.Request) {
	rr := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodGet, path, nil)
	return rr, rq
}

func setupTime(hours, minutes, seconds int) time.Time {
	return time.Date(1970, 1, 1, hours, minutes, seconds, 0, time.UTC)
}

func buildResponseFrom(rdr io.Reader) interface{} {
	type response struct {
		Hour   int
		Minute int
		Second int
	}
	var r response
	json.NewDecoder(rdr).Decode(&r)
	return r
}

func assertHasStatus(t *testing.T, rr *httptest.ResponseRecorder, want int) {
	t.Helper()
	got := rr.Code
	if want != got {
		t.Fatalf("Wanted status %d, got %d", want, got)
	}
}

func assertDeepEqual(t *testing.T, want interface{}, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted response %+v, got %+v", want, got)
	}
}

func assertHasContentType(t *testing.T, rr *httptest.ResponseRecorder, want string) {
	t.Helper()
	got := rr.Result().Header.Get("content-type")
	if want != got {
		t.Fatalf("Wanted content-type %s, got %s", want, got)
	}
}
