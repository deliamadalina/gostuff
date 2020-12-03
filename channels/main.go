package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"http://golang.org",
		"https://amazon.com",
		"https://amazonn.com",
	}

	c := make(chan string) // craetes a channel to comunicate between go main routine and go chiled routines

	for _, link := range links {
		go checkLink(link, c)
	}
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	for l := range c { // wait for the channel to return some value and assign it to 'l'
		//go checkLink(l, c) // spone a new go rutine passing the link and the channel
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)

		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might me down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
