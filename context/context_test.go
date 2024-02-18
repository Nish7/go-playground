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
	t         *testing.T
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *StubStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancelled")
	}
}

func (s *StubStore) assertWasNotCancelled() {
	s.t.Helper()

	if s.cancelled {
		s.t.Error("store was told to cancelled")
	}
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("return data from the store", func(t *testing.T) {
		data := "hello, world"
		store := &StubStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %s, wanted %s", response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {

		data := "hello, world"
		store := &StubStore{response: data, t: t}
		svr := Server(store)

		// mocks the request object to be sent to mock server
		request := httptest.NewRequest(http.MethodGet, "/", nil)

		// with cancel return a new context and cancel func, which invoked does finish the
		cancellingCtx, cancel := context.WithCancel(request.Context())

		request = request.WithContext(cancellingCtx)

		// this runs ayncrnous on a go routine
		time.AfterFunc(5*time.Millisecond, cancel)

		response := httptest.NewRecorder()
		//  mocks an response writer interface, so it can capture when handler writes to it, allowing
		// us to inspect and verify
		// this way of testing is dependecy injection

		svr.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}
