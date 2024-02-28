package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speed of server , returning the fastest one", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := WebsiteRacer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("timesout after 10 sec no response", func(t *testing.T) {

		slowServer := makeDelayedServer(11 * time.Millisecond)
		fastServer := makeDelayedServer(12 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		_, err := WebsiteRacer(slowURL, fastURL)

		if err == nil {
			t.Errorf("expected an error but dint get one")
		}

	})
}

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
}
