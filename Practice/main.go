package main

import (
	"fmt"
	"strings"
)

type Object struct {
	name string
	age  int
}

func main() {
	
	/// slice with append
	var arrays []string
	arrays = append(arrays, "X")
	arrays = append(arrays, "Y")
	arrays = append(arrays, "Z")
	fmt.Print(arrays)

	/// 2d slice with append
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board = append(board, []string{"X", "_", "_"})
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
