package main

import "fmt"

func main() {
	var numbers [5]int

	fmt.Println("Enter five numbers (1-30):")
	for i := 0; i < 5; i++ {
		fmt.Printf("Number %d: ", i+1)
		fmt.Scanln(&numbers[i])

		if numbers[i] < 1 || numbers[i] > 30 {
			fmt.Println("Invalid input! Enter a number between 1 and 30.")
			i--
		}
	}

	fmt.Println("\nBar Chart:")
	for _, num := range numbers {
		for i := 0; i < num; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
