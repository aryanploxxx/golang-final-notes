package main

import "fmt"

func activateGiftCard() func(int) int {
	amount := 100

	debitfunc := func(debitAmount int) int {
		amount += debitAmount
		return amount
	}

	return debitfunc
}

func main() {
	useGiftCard1 := activateGiftCard()
	useGiftCard2 := activateGiftCard()

	fmt.Println("Gift Card 1 Balance:", useGiftCard1(10))
	fmt.Println("Gift Card 2 Balance:", useGiftCard2(5))
}
