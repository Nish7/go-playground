package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// Channel Types:
// Unbuffered Channels: These channels do not store any values. Sending a value **blocks** the sending goroutine until another goroutine receives the value from the channel, and vice versa for receiving.
// Buffered Channels: Created with a buffer capacity. Sending is blocked only when the buffer is full. Similarly, receiving blocks only when the buffer is empty. They are initialized like ch := make(chan Type, capacity).
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	res := make(map[string]bool)
	channel := make(chan result, len(urls))

	for _, url := range urls {
		go func(url string) {
			// We are running a seperate process or threads for each channel input
			// Sending into an unbuffered channel, why unbuffered is because it defaults that if ommited
			// So.., if it is unbuffered, that means we it is technically **blocked**, until this 2nd loop
			// doesnt recieves
			// same for the reciever it will keep looking or **blocked** until it recieves a data from the channel
			// it is very much expecting
			channel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r, ok := <-channel

		if !ok {
			return nil
		}

		res[r.string] = r.bool
	}

	return res
}
