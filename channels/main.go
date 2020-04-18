package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://google.com",
	}

	// a channel is for communicate between go routines (threads)
	c := make(chan string)

	for _, link := range links {
		// go keyword to create a new routine (concurrency)
		// sends channel as a parameter to get message from that routine
		go checkLink(link, c)
	}

	for l := range c { // wait for the channel to return a value and assign to the variable 'l' and run loop body
		// waiting for channel message (blocking operation) and create new routine with that channel
		// pass the link and channel as a argument
		go func(link string) { //function literal (like lamble in c#, anonymous function like js)
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // sync request

	if err != nil {
		fmt.Println(link, "might be down!")
		// send this message (link) to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// send this message (link) to the channel
	c <- link
}
