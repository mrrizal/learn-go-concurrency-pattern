package inifiniteforloop

import (
	"fmt"
	"time"
)

func newsFeed(ch chan string) {
	defer close(ch)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Second / 2)
		ch <- fmt.Sprintf("News: %d", i)
	}
}

func printAllNews(ch chan string) {
	for {
		select {
		case msg := <-ch:
			// channel will send emtpy value when channel is closed
			// so, if channel is closed, we just return it
			if msg == "" {
				return
			}
			fmt.Println(msg)
		//	if there is no message anymore after 5 second, select statement will execute this
		case <- time.After(5 * time.Second):
			return
		}
	}
}

func Example() {
	news := make(chan string)
	go newsFeed(news)
	printAllNews(news)
}
