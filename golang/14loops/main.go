package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println("Days of the week: ", days)
	// Days of the week:  [Sunday Tuesday Wednesday Friday Saturday]

	// This method is not as useful in scenarios we know that we need to iterate through the entire list
	// TYPE - 1
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	// len(days) would have given fixed value had we taken array instead of slice
	/*
		Sunday
		Tuesday
		Wednesday
		Friday
		Saturday
	*/

	// TYPE - 2
	for i := range days {
		fmt.Println(days[i])
		// i here returns us the index of the slice
	}

	// TYPE - 2
	for index, day := range days {
		fmt.Printf("Index: %v, Day: %v\n", index, day)
	}
	// Index: 0, Day: Sunday
	// We get 2 values from the range keyword, the index and the value at that index
	// index and day are 2 types of variables which return the index and the value at that index
	// This is similar to the for each loop in other languages

	// TYPE - 3
	rougueValue := 1
	for rougueValue < 5 {
		if rougueValue == 3 {
			goto bookmark
		}
		if rougueValue == 5 {
			break
		}
		fmt.Println("Value of rogueValue is ", rougueValue)
		rougueValue++
	}
	// Similar to while loop in other languages

bookmark:
	fmt.Println("This is a bookmark")
}
