package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter user input: ")

	input, _ := reader.ReadString('\n')

	fmt.Println("You rated us: ", input)

	// Now for some reason we want to add one to the rating user entered.
	// We would need to conver the string to an integer
	// We can use the strconv package to convert the string to an integer
	/*
		numRating, err := strconv.ParseFloat(input, 64)
		// strconv.ParseFloat() takes two arguments, the first is the string we want to convert and the second is the bit size, i.e 32 or 64 to decide on the precision
		-> This will give the following error:
			strconv.ParseFloat: parsing "4\r\n": invalid syntax
			-> This is because the input has a new line character at the end
			-> We will have to trim the input to get rid of nay extra spaces or characters
	*/

	// strconv has several options to convert strings to integers, floats, etc. like ParseComplex, ParseInt, ParseFloat, etc.

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You rated us: ", numRating+1)
	}
}
