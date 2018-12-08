package main

import "fmt"

type example struct {
	pi      float32
	counter int16
	flag    bool
}

func main() {
	var e1 example
	fmt.Printf("e1: [%v]\n", e1)
	e2 := example{
		pi:      3.14,
		counter: 0,
		flag:    true,
	}
	fmt.Printf("e2: [%v]\n", e2)
}
