package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	// os.Stdin is a file object that represents the standard input stream and tells NewReader where should we be reading from
	// reader here acts like a pointer, a reference that this is what we are supposed to do
	fmt.Println("Enter rating for our Pizza: ")
	// a new line would automatically get inserted  because of Println, so we don't need to add \n

	// Now we want to store the value whatever the reader reads(user enters)
	// This is where we use comma, ok or comma, error syntax -> it basically is the alternative of try-catch in Go
	// We are trying to read the input from the user and if there is an error, we are catching it

	input, _ := reader.ReadString('\n')
	// \n indicates that we want to read the input until the user presses the enter key
	// We can specify reading different formats - ReadByte, ReadBytes, ReadS1ice, ReadString, ReadLine, ReadRune, UnreadByte, UnreadRune
	// input, err -> first part is the input and second is the error, if the error is not useful to use, we will simply use _, if input is not useful, we will use _ for that like  _, err

	fmt.Println("Thank you for rating us: ", input)

	fmt.Printf("Thank you for rating us: %T", input)
	// It will return 'string' as the rating type, because ReadString returns a string
	// This will cause problems as we will only be able to use string related functions on it
}
