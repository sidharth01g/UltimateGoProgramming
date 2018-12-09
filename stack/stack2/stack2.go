package main

const Size = 512
const Maxdepth = 5000
// var s_address *string

func stackCopy(s *string, recursion_depth int, dummy [Size]int16) {
	recursion_depth++
	println("Depth: ", recursion_depth, " s: ", s, )
	// if s != s_address {
	// 	println("Stack reallocated")
	// 	s_address = s
	// }
	if recursion_depth > Maxdepth {
		return
	}
	stackCopy(s, recursion_depth, dummy)
}


func main() {
	s := "12345"
	recursion_depth := 0
	// s_address = &s

	// Initial stack for Go: 2 KB
	dummy := [Size]int16{} // Size * 2 bytes = 512 * 2 bytes = 1 KB
	stackCopy(&s, recursion_depth, dummy)
}