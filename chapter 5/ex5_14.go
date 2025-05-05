package main

import (
	"fmt"
)

func main() {
	principal := 100000
	years := 10

	fmt.Println("Rate | Final Amount ($)")
	fmt.Println("------------------------")

	for rate := 5; rate <= 10; rate++ {
		finalAmount := principal

		for year := 1; year <= years; year++ {
			finalAmount = finalAmount + (finalAmount * rate / 100)
		}

		dollars := finalAmount / 100
		cents := finalAmount % 100
		fmt.Printf("%d%%  | $%d.%02d\n", rate, dollars, cents)
	}
}
