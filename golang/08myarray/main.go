package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitList [4]string
	//  It is a compulsion in arrays to specify the length of the array while creating it
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Banana"
	/*
		var fruitList [4]int
		fruitList[0] = 1
		fruitList[1] = 2
		fruitList[3] = 4
		fmt.Println("Fruit list is:", fruitList)
		-> Fruit list is: [1 1 0 1]
		-> It replaces empty values by "" in string data type and 0 in integer data type
	*/

	fmt.Println("Fruit list is:", fruitList)
	// Fruit list is: [Apple Tomato  Banana]
	// Notice that the element at index 2 is empty

	fmt.Println("Length of Fruit list is:", len(fruitList))
	// Fruit list is: 4
	// No matter how many elements are initialized, the length of the array will be 4 as it is what you specified while creating the array

	// Another way to declare an array
	// vegList := [3]string{"Potato", "Beans", "Brinjal"}

	var vegList = [5]string{"potato", "beans", "brinjal"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Printf("Type of Veg list is: %T", vegList)
	// [5]string
	fmt.Println("Length of Veg list is:", len(vegList))
	// 5
}
