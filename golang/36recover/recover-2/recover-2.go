package main

import (
	"fmt"
	"time"
)

func main() {
	a := []string{"a", "b"}
	checkAndPrintWithRecover(a, 2)
	time.Sleep(time.Second)
	fmt.Println("Exiting normally")
}

func checkAndPrintWithRecover(a []string, index int) {
	defer handleOutOfBounds()
	go checkAndPrint(a, index)
}

func checkAndPrint(a []string, index int) {
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

// An important point to note about be recover function is that it can only recover the panic happening in the same goroutine.  If the panic is happening in different  goroutine and recover is in a different goroutine then it won’t stop  panic. Lets see a program for that
