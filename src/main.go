package main

import (
	"fmt"
)

func change(a []int) {
	a = append(a, 1)
	fmt.Printf("%p\n", a)
	fmt.Println(a)
}
func main() {
	s := make([]int, 5)
	fmt.Println(s)
	change(s)
	fmt.Println(s)
	fmt.Printf("%p\n", s)
}
