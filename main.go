package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Ready to Go!")
	three := add(1, 2)
	fmt.Println(three)
}
