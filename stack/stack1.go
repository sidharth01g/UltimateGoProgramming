package main

// import "fmt"

func main() {
	count := 10
	println("Before: ", count, &count)
	count++
	println("After", count, &count)
}