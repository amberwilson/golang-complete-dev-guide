package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}
	fmt.Println(<-channel)
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		channel <- "Might be down"
		return
	}

	fmt.Println(link, "is up")
	channel <- "It is up"
}
