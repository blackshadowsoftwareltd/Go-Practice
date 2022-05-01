package main

import (
	"fmt"
)

func main() {
	firstBill := newBill("John")

	firstBill.addItem("onion soup", 4.45)
	firstBill.addItem("veg pie", 8.50)
	firstBill.addItem("toffee pudding", 4.95)
	firstBill.addItem("coffee", 2.25)

	firstBill.updateQuantity(4)

	fmt.Println(firstBill.format()) //? call using bill struct (second way)

}
