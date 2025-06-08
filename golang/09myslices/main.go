package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")
	// Slices are under the hood arrays and they will be more useful as they are more powerful

	// In slices, we do not need to specify the size of the array
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	// Although in this method of defining, we have to initialize the slice
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	// Type of fruitList is []string

	// Since arrays are fixed in size, we can't append to them
	// But we can append to slices and it automatically resizes and expands the memory for us

	fruitList = append(fruitList, "Banana")
	fmt.Println("Updated Fruit List is: ", fruitList)
	// Apple Tomato Peach Banana

	fruitList = append(fruitList[1:]) // This will remove the first element of the slice i.e. the 0th index
	fmt.Println("Updated Fruit List is: ", fruitList)
	// Tomato Peach Banana

	fruitList = append(fruitList[1:3]) // Here, the original fruitlist after appending banana is being considered.
	// 0 1 2 3 4
	// The above are indexes in the original fruitlist, it will slice the slice and return a small part of the original slice i.e. indexes 1 and 2
	// The first index is inclusive and the second index is exclusive
	// So the above will return the elements at index 1 and 2, 3 will not be included
	fmt.Println("Updated Fruit List is: ", fruitList)
	// Peach Banana

	fruitList = append(fruitList[:3])
	// Starts from default value of the first index which is 0
	// This will return the elements from index 0 to 2, 3 will not be included
	fmt.Println("Updated Fruit List is: ", fruitList)

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++

	highScores := make([]int, 4)
	// (dataype, length)
	// THis will by default create an slice which is an abstraction over arrays
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // This will give an error as the index is out of range as we had initialized the slice with a size of 4

	highScores = append(highScores, 555, 666, 777)
	// Interestingly, this method will work
	// What will happen is that 'make' did a default allocation of memory for 4 elements
	// But, append method will resize the slice and reallocate the memory and new values are accommodated
	fmt.Println("High Scores: ", highScores)

	sort.Ints(highScores)
	// sort is a package in Go which has a method Ints which will sort the slice of integers in ascending order
	// 234 465 555 666 777 867 945
	fmt.Println("Sorted High Scores: ", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	// This will return a boolean value if the slice is sorted or not
	// true

	// ++++++++++++++++++++++++++++++++++++++++++++++++++++

	// # How to remove a value from slice based on index

	var courses = []string{"react", "javascript", "swift", "python", "ruby"}
	fmt.Println("Courses: ", courses)
	var index int = 2
	// This is the index which we index to delete
	/*
		courses = append(courses[:index], courses[index+1:])
		-> This feels logically correct as we are just adding up two parts of the array into one, but is incorrect
	*/
	courses = append(courses[:index], courses[index+1:]...)
	// react javascript python ruby
	fmt.Println("Updated Courses: ", courses)
}
