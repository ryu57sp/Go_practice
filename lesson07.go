package main

import (
	"fmt"
)

func main() {
	//要素数省略
	//a := [...]string{"sato", "suzuki", "takahashi"}
	a := [3]string{"sato", "suzuki", "takahashi"}
	a[1] = "tanaka"

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
}
