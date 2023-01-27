package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"https://www.ebay.com/", "https://github.com/", "https://www.php.net/"}

	Fetch(urls)
}

func Fetch(urls []string) {
	successful := true
	responses := make(chan *http.Response)

	for _, url := range urls {
		go func(url string) {
			client := http.Client{
				Timeout: 1 * time.Second,
			}

			resp, err := client.Get(url)

			if err != nil {
				successful = false
				fmt.Println(err)
				return
			}

			defer resp.Body.Close()

			responses <- resp
		}(url)
	}

	if successful {
		for i := 0; i < len(urls); i++ {
			resp := <-responses
			if resp.StatusCode == 200 {
				fmt.Println(resp.Status)
			}
		}
	}
}
