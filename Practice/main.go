package main

import (
	"fmt"
	"sort"
)

func main() {
	var users []User = []User{
		{"John", 20},
		{"Jane", 21},
		{"Jill", 23},
		{"Joe", 24},
		{"Jack", 22},
	}
	fmt.Println(users)

	///? sorting by age (int)
	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	fmt.Println(users)

	///? sorting by name (string)
	sort.Slice(users, func(x, y int) bool {
		return users[x].Name < users[y].Name
	})
	fmt.Println(users)

}

type User struct {
	Name string
	Age  int
}
