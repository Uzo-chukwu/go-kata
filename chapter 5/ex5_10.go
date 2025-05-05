package main

import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0
	years := 10

	fmt.Println("Rate | Final Amount")
	fmt.Println("-------------------")

	for rate := 5.0; rate <= 10.0; rate++ {
		finalAmount := principal * math.Pow(1+(rate/100), float64(years))
		fmt.Printf("%.0f%%  | $%.2f\n", rate, finalAmount)
	}
}
