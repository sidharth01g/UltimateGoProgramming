package main

import "fmt"

func main() {
	var a int
	a = 1
	b := 10

	fmt.Printf("a: %T, [%v]\n", a, a)
	fmt.Printf("b: %T, [%v]\n", b, b)
}