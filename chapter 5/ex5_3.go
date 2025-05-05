package main

import "fmt"

func main() {
	sum := 0

	for count := 1; count <= 99; count += 2 {
		sum += count
	}

	fmt.Println("Sum of odd integers:", sum)
}
