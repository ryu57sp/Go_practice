package main

import "fmt"

type User struct {
	gender string
	age    int
}

func main() {
	u := User{gender: "male", age: 20}

	fmt.Println(u)
}

// func main() {
// 	var u User

// 	u.gender = "male"
// 	u.age = 20

// 	fmt.Println(u)
// }
