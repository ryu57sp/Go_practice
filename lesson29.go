package main

import "fmt"

type Student struct {
	name          string
	math, english float64
}

func main() {
	//s := Student{80, "sato", 70}　順番逆にできない
	//s := Student{"sato", , 70} 空欄にできない
	//s := Student{"sato", 80, 70} OK

	s := Student{name: "sato", math: 80, english: 70}

	fmt.Println(s)
}
