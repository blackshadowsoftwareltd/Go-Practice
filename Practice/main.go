package main

import (
	"fmt"
	"strings"
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

	///? searching in slice by name (raw way)
	var newUsers []User
	for _, u := range users {
		if strings.Contains(strings.ToLower(u.Name), strings.ToLower("k")) {
			newUsers = append(newUsers, u)
		}
	}
	fmt.Println(newUsers)
	///? is contsins in a slice
	fmt.Println("is contains", isContainsInSlice(users, "x"))
}

///? is contains in a slice
func isContainsInSlice(s []User, name string) bool {
	for _, u := range s {
		if strings.Contains(strings.ToLower(u.Name), strings.ToLower(name)) {
			return true
		}
	}
	return false

}

type User struct {
	Name string
	Age  int
}
