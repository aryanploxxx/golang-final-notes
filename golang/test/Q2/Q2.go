package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "orange"}
	fmt.Println("Fruits slice currently has: ", fruits)
	fmt.Println("Options: \n1. Add \n2. Delete \n3. Update \n4. Exit")

	for {
		fmt.Println("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter the fruit you want to add:")
			var newFruit string
			fmt.Scanln(&newFruit)
			fruits = append(fruits, newFruit)
			fmt.Println("Updated Fruits slice currently has: ", fruits)

		case 2:
			fmt.Println("Enter the fruit index you want to remove:")
			var choice int
			fmt.Scanln(&choice)
			fruits = append(fruits[:choice], fruits[choice+1:]...)
			fmt.Println("Updated Fruits slice currently has: ", fruits)

		case 3:
			fmt.Println("Enter the fruit index you want to update:")
			var choice int
			fmt.Scanln(&choice)
			if choice < 0 || choice >= len(fruits) {
				fmt.Println("Invalid index")
				break
			}
			fmt.Println("Enter the name of the newfruit:")
			var newFruit string
			fruits[choice] = newFruit
			fmt.Println("Updated Fruits slice currently has: ", fruits)

		case 4:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
