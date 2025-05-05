package main

import (
	"fmt"
)

func main() {
	var number, largest int

	fmt.Println("Enter 10 numbers:")

	for counter := 1; counter <= 10; counter++ {
		fmt.Printf("Enter number %d: ", counter)
		fmt.Scanln(&number)

		if number > largest {
			largest = number
		}
	}

	fmt.Printf("\nThe largest number entered is: %d\n", largest)
}
