package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job, wg *sync.WaitGroup) {
	fileName := j.text + ".txt"
	fileContents := ""
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fileContents += j.text
		fmt.Println(j.text)
		j.i++
	}
	err := ioutil.WriteFile(fileName, []byte(fileContents), 0644)
	if err != nil {
		panic("Something went awry")
	}

	wg.Done()

}

func main() {

	wg := new(sync.WaitGroup)

	hello := new(Job)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	world := new(Job)
	world.text = "world"
	world.i = 0
	world.max = 5

	go outputText(hello, wg)
	go outputText(world, wg)

	wg.Add(2)
	wg.Wait()
}
