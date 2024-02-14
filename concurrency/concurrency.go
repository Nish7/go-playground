package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	res := make(map[string]bool)

	for _, url := range urls {
		res[url] = wc(url)
	}

	return res
}
