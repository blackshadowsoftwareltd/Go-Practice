package main

import (
	"fmt"
)

func main() {
	names:=[]string {"Bob","John","Mike","Tom"}
	fmt.Println(names)
	temp:=names[1:3]
	fmt.Println(temp)
	temp[0]="MR. X"
	fmt.Println(temp)
	fmt.Println(names)
 }

 
