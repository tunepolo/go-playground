package main

import (
	"fmt"
	"log"
	"net/http"
)

var empty struct{}

func getStatus(urls []string) <-chan string {
	statusChannel := make(chan string, 3)
	limit := make(chan struct{}, 5) // goroutineの同時実行数を5に制限

	go func() {
		for _, url := range urls {
			select {
			case limit <- empty:
				go func(url string) {
					res, err := http.Get(url)
					if err != nil {
						log.Fatal(err)
					}
					defer res.Body.Close()
					statusChannel <- fmt.Sprintf("%s -> %d", url, res.StatusCode)

					<-limit
				}(url)
			}
		}
	}()

	return statusChannel
}

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.yahoo.co.jp",
		"http://github.com",
	}

	statusChannel := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChannel)
	}
}
