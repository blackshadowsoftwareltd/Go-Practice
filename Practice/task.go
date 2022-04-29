package main

import (
	"fmt"
	"sort"
)

func doSomeThing() {
	_list := []int{5, 6, 7, 1, 2, 3, 10, 4, 8, 9}
	fmt.Println(_list)

	sort.Ints(_list) //? sort

	fmt.Println(_list)

	fmt.Println(sort.SearchInts(_list, 10)) //? index of list's index
}
