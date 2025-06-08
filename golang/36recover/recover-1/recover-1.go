package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	a := []string{"a", "b"}
	checkAndPrintWithRecover(a, 2)
	fmt.Println("Exiting normally")
}

func checkAndPrintWithRecover(a []string, index int) {
	defer handleOutOfBounds()
	checkAndPrint(a, index)
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
		fmt.Println("Stack Trace:")
		debug.PrintStack()
	}
}

// The above program is quite the same as the previous program other than we have an additional function checkAndPrintWithRecover which contains the call to

// defer function with recover which is handleOutOfBounds
// calls checkAndPrint function
// So basically checkAndPrint function raises the panic but doesn’t have the recover function instead call to recover lies in the checkAndPrintWithRecover function. But still the program is able to recover from panic  as panic can also be recovered in the called function also and subsequently in the chain as well

// In the above program we print the stack trace of the panic in the recover function using the StackTrace function. It prints the correct  stack trace as seen from the output
// Debug package of golang also provide StackTrace function that can be used print the stack trace of the panic in the recover function
