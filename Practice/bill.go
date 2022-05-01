package main

import "fmt"

type bill struct {
	buyerName string
	quantity  int
	items     map[string]float64
}

func newBill(name string) bill {

	b := bill{buyerName: name, quantity: 1, items: map[string]float64{"pie": 2.99, "cake": 3.99}}
	return b

}

func (b bill) format() string {
	fs := "Bill breakdown : \n"
	var total float64 = 0

	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v $%v \n", key+" :", value) //? Sprint & Sprintf
		total += value
	}
	fs += fmt.Sprintf("%-25v $%0.2f", "total :", total) //? %0.2f means 2 charactar after the decimal point
	//?-25 means, it takes 25 character space after this value. ex (%-25v)
	//? only 25 takes 25 character space before this value. ex (%25v)
	return fs
}

func (b bill) updateQuantity(quantity int) {
	b.quantity = quantity
}
func (b bill) addItem(name string, price float64) { b.items[name] = price }
