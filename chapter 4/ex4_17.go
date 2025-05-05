package main

import (
	"fmt"
)

func main() {
	var name string
	var earnings, tax float64

	for i := 1; i <= 3; i++ {
		fmt.Printf("Enter name of citizen %d: ", i)
		fmt.Scanln(&name)

		fmt.Printf("Enter %s's yearly earnings: ", name)
		fmt.Scanln(&earnings)

		if earnings <= 30000 {
			tax = earnings * 0.15
		} else {
			tax = (30000 * 0.15) + ((earnings - 30000) * 0.20)
		}

		fmt.Printf("%s's total tax: $%.2f\n\n", name, tax)
	}
}
