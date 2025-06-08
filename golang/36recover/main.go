package main

import "fmt"

func main() {

	a := []string{"a", "b"}
	checkAndPrint(a, 2)
	fmt.Println("Exiting normally")
}

func checkAndPrint(a []string, index int) {
	defer handleOutOfBounds()
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

// The recover function returns the value which was passed to the panic function. Therefore it is a good practice to check the return value of the recover function. If the return value is nil then panic did not happen and recover function was not called with the panic. That is why we have below code in the  defer function handleOutofBounds
// Here if r is nil then panic did not happened. So if there is no panic then call to recover will return nil
