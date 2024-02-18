package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response  string
	cancelled bool
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "hello, world"
	svr := Server(&StubStore{data, false})

	request := httptest.NewRequest(http.MethodGet, "/", nil) // mocks the request object to be sent to mock server

	cancellingCtx, cancel = context.WithCancel(request.Context())
	time.AfterFunc(5*time.Millisecond, cancel)

	response := httptest.NewRecorder() //  mocks an response writer interface, so it can capture when handler writes to it, allowing us to inspect anad verify
	// this way of testing is dependecy injection

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("got %s, wanted %s", response.Body.String(), data)
	}
}
