package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"http://www.yahoo.co.jp",
		"http://github.com",
	}

	wait := new(sync.WaitGroup)

	for _, url := range urls {
		wait.Add(1)

		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.StatusCode)

			wait.Done()
		}(url)
	}

	wait.Wait()
}
