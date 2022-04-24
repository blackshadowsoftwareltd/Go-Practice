package main

import (
	"fmt"
	"go/types"
)

func main(){
	var i interface{} = "hello"

	s := i.(types)
	fmt.Println(s)
}