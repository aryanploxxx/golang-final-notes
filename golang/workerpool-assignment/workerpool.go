package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"sync"
	"time"
)

var profile = flag.String("cpuprofile", "", "output pprof data to file")

// go run pprof.go -cpuprofile=cpu.prof
// go tool pprof cpu.prof

var tracefile = flag.String("tracefile", "", "output trace data to file")

// go run workerpool.go -tracefile trace.out
// go tool trace trace.out

func main() {

	flag.Parse()
	if *profile != "" {
		flag, err := os.Create(*profile)
		if err != nil {
			fmt.Println("Could not create profile", err)
		}
		pprof.StartCPUProfile(flag)
		defer pprof.StopCPUProfile()
	}

	// Start tracing if enabled
	if *tracefile != "" {
		traceFile, err := os.Create(*tracefile)
		if err != nil {
			fmt.Println("Could not create trace file:", err)
			return
		}
		defer traceFile.Close()
		if err := trace.Start(traceFile); err != nil {
			fmt.Println("Could not start trace:", err)
			return
		}
		defer trace.Stop()
	}

	var tasks []string = []string{"task1", "task2", "task3", "task4", "task5", "task6", "task7", "task8", "task9", "task10", "task11", "task12", "task13", "task14", "task15", "task16", "task17", "task18", "task19", "task20", "task21", "task22", "task23", "task24", "task25"}
	var workers []string = []string{"worker1", "worker2", "worker3"}

	var tasksChann chan string = make(chan string, len(tasks))
	var workersChann chan string = make(chan string, len(workers))

	fmt.Println("Worker Pool Started")

	var wg sync.WaitGroup

	go func() {
		for _, task := range tasks {
			tasksChann <- task
		}
	}()

	go func() {
		for _, worker := range workers {
			workersChann <- worker
		}
	}()

	for i := 0; i < 20; i++ {
		task := <-tasksChann
		worker := <-workersChann
		wg.Add(1)
		go executeTask(task, worker, &wg, workersChann)
	}

	wg.Wait()

	fmt.Println("Reacher here")

	if len(tasks) == 0 {
		for worker := range workersChann {
			fmt.Println("Worker", worker, "is free")
		}
	}

	close(tasksChann)
	close(workersChann)

	if len(tasks) == 0 {
		fmt.Println("All tasks completed!")
	}
}

func executeTask(task string, worker string, wg *sync.WaitGroup, workersChann chan string) {
	defer wg.Done()
	fmt.Println("Executing task", task, "by worker", worker)
	time.Sleep(2 * time.Second)
	fmt.Println(task, "completed by worker", worker)
	workersChann <- worker
}
