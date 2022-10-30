package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url    string
	isGood bool
}

func CheckWebsittes(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{url: u, isGood: wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.isGood
	}

	return results
}
