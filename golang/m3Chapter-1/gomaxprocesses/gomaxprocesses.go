package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

// func listThreads() int {
// 	threads := runtime.GOMAXPROCS(0)
// 	return threads
// }

func showNumber(num int) {
	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num, tstamp)
	time.Sleep(time.Millisecond * 10)
}

func main() {
	// fmt.Printf("%d thread(s) available to Go.", listThreads())
	runtime.GOMAXPROCS(0)
	iterations := 10
	for i := 0; i <= iterations; i++ {
		go showNumber(i)
	}
	runtime.Gosched()
}
