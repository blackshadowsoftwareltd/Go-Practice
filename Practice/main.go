package main

import "fmt"

func main() {
 s:="Hello World"
 found(s)
 i:=10
 found(i)
}
var emptyInterface interface{}

func found(i interface{}){
	fmt.Printf("Type %T, value %v\n",i,i)
}