package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	myBill := createBill()
	promptOptions(myBill)
	fmt.Println(myBill)
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name : ", reader)

	b := newBill(name)
	fmt.Println("Created a new bill - ", b.buyerName)

	return b
}

func getInput(prompt string, reader *bufio.Reader) (string, error) {

	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Chose option a - add item, s -save bill, q - add quantity", reader)
	fmt.Println(opt)
}
