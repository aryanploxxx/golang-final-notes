package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go goOne(ch1)
	go goTwo(ch2)
	select {

	case msg1 := <-ch1:
		fmt.Println(msg1)
	case ch2 <- "To goTwo goroutine":
	}
}

func goOne(ch chan string) {
	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}
