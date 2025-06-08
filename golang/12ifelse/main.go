package main

import "fmt"

func main() {
	fmt.Println("If Else")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 Login Count"
	}

	fmt.Println("Result is: ", result)

	// We dont have to necessarily create variables beforehand, we can create them on the go
	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	// We can also create variables in the if else block and checking it on the go
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	if err != nil {
		fmt.Println(err)
	}

}

/*
	if loginCount < 10 {
		result = "Regular User"
	}
	-> this is the correct way of writing curly braces in if else

	if loginCount < 10
	{
		result = "Regular User"
	}
	-> this way  will result in error
*/
