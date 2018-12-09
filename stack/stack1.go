package main

// import "fmt"
func increment_bad(val int) {
	val ++
	println("increment_bad - val: ", val, &val)
}

func increment_good(val *int) {
	*val++
	println("increment_good - val: ", *val, val, &val)
}

func main() {
	count := 10
	println("main - Before: ", count, &count)
	count++
	println("main - After", count, &count)
	increment_bad(count)
	println("main - After", count, &count)
	increment_good(&count)
	println("main - After", count, &count)
}