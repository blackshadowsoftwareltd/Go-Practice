package main

import (
	"fmt"
)

func main() {
	firstBill := newBill("John")
	fmt.Println(firstBill)

	// fmt.Println(bill.format(firstBill)) //? call using bill struct (first way)

	fmt.Println(firstBill.format()) //? call using bill struct (second way)

}
