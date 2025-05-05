package main

import (
	"fmt"
)

func main() {
	var currentPopulation float64 = 8.2e9
	var growthRate float64 = 0.0085
	var year int = 1
	var doubledYear int = 0
	var doublePopulation float64 = currentPopulation * 2

	fmt.Println("Year | Population | Increase")
	fmt.Println("--------------------------------")

	for year <= 75 {
		increase := currentPopulation * growthRate
		currentPopulation += increase

		fmt.Printf("%4d | %.0f | %.0f\n", year, currentPopulation, increase)

		if doubledYear == 0 && currentPopulation >= doublePopulation {
			doubledYear = year
		}

		year++
	}

	fmt.Printf("\nThe population will double in approximately year %d.\n", doubledYear)
}
