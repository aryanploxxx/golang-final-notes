package main

import (
	"fmt"
	"runtime"
	"strings"
)

var loremIpsum string
var finalIpsum string

// var letterSentChan chan string

func deliverToFinal(letter string, finalIpsum *string) {
	*finalIpsum += letter
}

func capitalize(current *int, length int, letters []byte, finalIpsum *string) {
	for *current < length {
		thisLetter := strings.ToUpper(string(letters[*current]))
		deliverToFinal(thisLetter, finalIpsum)
		*current++
	}
}

// Problem: The shared variable *current is accessed and modified by multiple goroutines without synchronization, leading to data races.
func main() {
	runtime.GOMAXPROCS(2)

	index := new(int)
	*index = 0

	loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscingelit. Vestibulum venenatis magna eget libero tincidunt, accondimentum enim auctor. Integer mauris arcu, dignissim sit ametconvallis vitae, ornare vel odio. Phasellus in lectus risus. Utsodales vehicula ligula eu ultricies. Fusce vulputate fringillaeros at congue. Nulla tempor neque enim, non malesuada arculaoreet quis. Aliquam eget magna metus. Vivamus laciniavenenatis dolor, blandit faucibus mi iaculis quis. Vestibulumsit amet feugiat ante, eu porta justo."

	letters := []byte(loremIpsum)
	// Converts loremIpsum to a byte slice (letters): Allows character-by-character manipulation.

	length := len(letters)

	go capitalize(index, length, letters, &finalIpsum)

	go func() {
		go capitalize(index, length, letters, &finalIpsum)
	}()

	/*
		go capitalize(index, length, letters, &finalIpsum)
		go capitalize(index, length, letters, &finalIpsum)
	*/

	/*
		Both goroutines share:
		The index pointer (shared position tracker).
		The finalIpsum pointer (shared result string).
		Both goroutines access and modify *index and *finalIpsum simultaneously without synchronization.
	*/

	fmt.Println(length, " characters.")
	fmt.Println(loremIpsum)
	fmt.Println(*index)
	fmt.Println(finalIpsum)

	// The final output in finalIpsum may contain missing or duplicated characters.
	// The value of *index may not equal length.

}
