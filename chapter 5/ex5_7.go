package main

import (
	"fmt"
)

func main() {
	var count, num, min, max int

	fmt.Print("Enter how many numbers: ")
	fmt.Scanln(&n)

	if count <= 0 {
		fmt.Println("Invalid input, must be at least 1.")
		return
	}

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num)
	min, max = num, num

	for i := 1; i < n; i++ {
		fmt.Print("Enter next number: ")
		fmt.Scanln(&num)

		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	fmt.Printf("Minimum: %d, Maximum: %d\n", min, max)
	fmt.Printf("Sum of extremes: %d\n", min+max)
}
