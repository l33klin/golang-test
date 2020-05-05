package main

import (
	"fmt"
	"math/rand"
	"time"
)

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		fmt.Println("Get a result")
		results[result.string] = result.bool
	}

	return results
}

func MyWebsiteChecker(w string) bool {
	rand.Seed(time.Now().UnixNano())
	wait_second := rand.Int31n(10)
	fmt.Printf("Wait %d seconds, %v\n", wait_second, w)
	time.Sleep(time.Duration(wait_second) * time.Second)
	return true
}

func main() {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	results := CheckWebsites(MyWebsiteChecker, websites)
	fmt.Println(results)
}
