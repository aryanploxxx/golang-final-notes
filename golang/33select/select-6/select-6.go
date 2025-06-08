package main

import (
	"fmt"
	"time"
)

// Select statement with a nil channel

func main() {
	news := make(chan string)
	go newsFeed(news)
	printAllNews(news)
}

func printAllNews(news chan string) {
	for {
		select {
		case n := <-news:
			fmt.Println(n)
			news = nil
		case <-time.After(time.Second * 1):
			fmt.Println("Timeout: News feed finished")
			return
		}
	}
}

func newsFeed(ch chan string) {
	for i := 0; i < 2; i++ {
		time.Sleep(time.Millisecond * 400)
		ch <- fmt.Sprintf("News: %d", i+1)
	}
}

// Send or receive operation on nil channel blocks forever. Hence a use case of having a nil channel in the select statement is to disable that case statement after the the send or receive operation is completed on that case statement. The channel then can simply be set to nil. That case statement will be  ignored when the select statement is executed again and receive or send operation will be waited on another case statement. So it is meant to ignore that case statement and execute the other case statement
// The above program is pretty much similar to the program we studied above related to having a select statement inside infinite for loop. The only change being that  after we receive the first news, we disabled the case statement by setting the news channel to nil.
