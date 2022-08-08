package main

import (
	"fmt"
)

func main() {
	name := "pashkan"
	fmt.Printf("address of b_ptr: %d\n", name)
	fmt.Printf("address of b_ptr: %d\n", &name)
}
