package main

import (
	"fmt"
	"time"
)

// Select statement with an infinite for loop outside

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

// In the above program, we have created a channel named news  which will hold data of string type. Then we pass this channel to the newsfeed function which is pushing the news feed to this channel . In the select statement, we are receiving the news feed from the news channel. This select statement is inside an infinite for loop  so the select statement will be executed multiple times until we  exit out of for loop . We also have time.After with a duration for 1 second as one of the case statements. So this set up will receive all the news from the news channel for  1  second and then exit.
