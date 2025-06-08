package main

import "fmt"

func main() {
	ch1 := make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	}

	go func() {
		ch1 <- "Hello"
	}()
}

// in this program, select statement will block indefinately and the below go routine will not get a chance to execute
