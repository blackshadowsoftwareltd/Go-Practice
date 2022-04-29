package main

import (
	"fmt"
	"sort"
)

func main() {
	_list := []int{5, 6, 7, 1, 2, 3, 10, 4, 8, 9}
	_datas := []string{"g", "h", "a", "b", "e", "f", "c", "j", "d", "i"}
	fmt.Println(_list)
	fmt.Println(_datas)

	sort.Ints(_list)     //? sort
	sort.Strings(_datas) //? sort

	fmt.Println(_list)
	fmt.Println(_datas)

	fmt.Println(sort.SearchInts(_list, 10))      //? index of list's index
	fmt.Println(sort.SearchStrings(_datas, "d")) //? index of list's index
}
