package main

import (
	"fmt"
	"sync"
)

func printhello(wg *sync.WaitGroup) {
	fmt.Println("hell oworld")
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)

	go printhello(&wg)
	go printhello(&wg)

	wg.Wait()

}
