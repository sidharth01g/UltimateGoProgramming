package main

type user struct {
	name string
	email string
}

func stayOnStack() user {
	u := user {
		name: "Alice",
		email: "alice@abc.com",
	}
	return u
}

// Memory will be allocated for this function on the heap
// because this function seems to return a pointer to a variable that is
// created within the function. If this function were to be allocated memory on the stack the varable would be'
// immediately erased when the function returns and thus the pointer returned would point to an adrress that has been deallocated.
// The Go compiler automatically detects this possibility and allocates memory on the heap instead
func escapeToHeap() *user {
	u := user {
		name: "Bob",
		email: "bob@xyz.com",
	}
	return &u
}

func main() {
	a := stayOnStack()
	b := escapeToHeap()

	println("a: ", a.name, a.email, &a)
	println("b: ", b.name, b.email, &b)
}