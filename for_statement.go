package main

import "fmt"

func main() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	// Use "for" instead of "while"
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
