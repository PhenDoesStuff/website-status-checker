package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //  Channel to make sure main func doesn't end before child routines

	for _, link := range links {
		go checkLink(link, c) // 'go' starts a brand new Go Routine to run code specific inside of checkLink
	}

	for l := range c {
		time.Sleep(time.Second)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return 
	}

	fmt.Println(link, "is up!")
	c <- link
}  