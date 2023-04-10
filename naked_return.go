package main

import "fmt"

func add_and_subtract(a, b int) (sum, gap int) {
	sum = a + b
	gap = a - b
	return
}

func main() {
	x, y := add_and_subtract(4, 10)
	fmt.Println(x, y)
}
