package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	WordbyWordScan()
	LinebyLineScan()
}

func WordbyWordScan() {
	file, err := os.Open("./scanner/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func LinebyLineScan() {
	file, err := os.Open("./sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// -> Read a large file Word by Word in Go(Golang)
// When it comes to reading large files, obviously we don’t want to load the entire file in memory.
// bufio package in golang comes to the rescue when reading large files in Golang.

// This is an example
// to show how
// to read file
// line by line.

// This
// is
// an
// example
// to
// show
// how
// to
// read
// file
// line
// by
// line
// and
// word
// by
// word.

// Please note that in the above program we set scanner.Split(bufio.ScanWords) which helps us to read the file word by word. Note however bufio.Scanner has max buffer size 64*1024 bytes which means in case you file has any line greater than the size of 64*1024, then it will give the error
