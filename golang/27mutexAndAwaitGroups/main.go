package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	// in case of RWMutex, we use RLock() and RUnlock() for reading and Lock() and Unlock() for writing

	var score = []int{0}

	// func(){}() // this is a self executing function - IIFE

	wg.Add(3)

	// Important thing to notice is that the goroutines are not aware the existance of the other
	// TO solve this exact problem, channels were made

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(1)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// go func(wg *sync.WaitGroup, m *sync.RWMutex) {
	// 	fmt.Println("Three R")
	// 	mut.RLock()
	// 	fmt.Println(score)
	// 	mut.RUnlock()
	// 	wg.Done()
	// }(wg, mut)

	wg.Wait()
	fmt.Println(score)
	// order or output is not fixed/guaranteed
}

/*
var wg sync.WaitGroup
Description: This declares a wg variable as a value of type sync.WaitGroup.
Initialization: This is automatically initialized and ready to use; no need to use new or make.
Usage:
Directly call methods like wg.Add, wg.Done, and wg.Wait on it.
This is the most common way to declare a sync.WaitGroup.


var wg *sync.WaitGroup
Description: This declares a pointer to a sync.WaitGroup object, but it is not initialized.
Initialization: You must initialize this pointer before use:
Using new(sync.WaitGroup) to allocate a new sync.WaitGroup.
Or assigning it a reference to an already initialized sync.WaitGroup (e.g., wg = &existingWG).
Usage:
You need to dereference the pointer when calling methods on it (e.g., (*wg).Add(1)), though Go allows implicit dereferencing, so wg.Add(1) is valid.
Useful when passing the WaitGroup as a pointer to functions


3. var wg &sync.WaitGroup{}
Description: This syntax is invalid. The & operator is used to take the address of a variable or value, and {} is used to initialize a struct. You cannot combine them in a declaration like this.
Correct Alternative: Use var wg sync.WaitGroup or wg := &sync.WaitGroup{}.
Why Invalid?
sync.WaitGroup does not require explicit initialization like &sync.WaitGroup{}.
The sync.WaitGroup struct cannot be initialized using composite literals because it contains unexported fields.
*/
