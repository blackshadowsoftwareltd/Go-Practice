package main

import (
	"fmt"
	"strconv"
)

type UserData struct {
	name string
	age  int
}

func main() {

	userInfo := make(map[string]UserData)
	for i := 0; i < 10; i++ {
		key := "user" + strconv.FormatInt(int64(i), 10)
		userInfo[key] = UserData{name: "Mr.X", age: i}
	}
	fmt.Println(userInfo)

}
