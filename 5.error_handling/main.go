package main

import (
	"fmt"
	"net/http"
)

func main() {
	type Result struct {
		Error    error
		Url      string
		Response *http.Response
	}

	// if your goroutine can reproduce error,
	// it's commonly to wrap it into struct then return it with your result/response
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp, Url: url}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://google.com", "https://badhost"}
	results := checkStatus(done, urls...)

	for result := range results {
		if result.Error != nil {
			fmt.Printf("%s: %s\n", result.Url, result.Error)
			continue
		}
		fmt.Printf("%s: %d\n", result.Url, result.Response.StatusCode)
	}
}
