package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := writeToTempFile("Some text")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("Write to file succesful")
}

func writeToTempFile(text string) error {
	file, err := os.Open("temp.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	n, err := file.WriteString(text)
	if err != nil {
		return err
	}
	fmt.Printf("Number of bytes written: %d", n)
	return nil
}

// In the above program, in the writeToTempFile function, we are opening a file and then trying to write some content to the file. After we have written the contents of the file we close the file. It is possible that during the write operation it might result into an error and function will return without closing the file. Defer function helps to avoid these kinds of problems. Defer function is always executed before the surrounding function returns. Let’s rewrite the above program with a defer function here.
// In the above program, we do defer file.Close() after opening the file. This will make sure that closing of the file is executed even if the write to the file results into an error. Defer function makes sure that the file will be closed regardless of number of return statements in the function

/*
	package main
	import "fmt"
	func main() {
		defer test()
		fmt.Println("Executed in main")
	}
	func test() {
		fmt.Println("In Defer")
	}
	-> We can also call a custom function in defer. Let’s see an example for that
*/

/*
	-> Inline Function in Defer

	package main

	import "fmt"

	func main() {
		defer func() { fmt.Println("In inline defer") }()
		fmt.Println("Executed")
	}

*/
