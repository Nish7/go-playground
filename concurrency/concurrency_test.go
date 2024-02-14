package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "wat:nishil.html" {
		return false
	}

	return true
}

func slowStubWebChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"wat:nishil.html",
		"http://example.com",
	}

	want := map[string]bool{
		"http://google.com":  true,
		"wat:nishil.html":    false,
		"http://example.com": true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
