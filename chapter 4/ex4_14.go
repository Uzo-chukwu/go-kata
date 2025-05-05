package main

import "fmt"

func main() {
	var miles, gallons, totalMiles, totalGallons float64
	var tripCount int

	for {

		fmt.Print("Enter miles driven (-1 to stop): ")
		fmt.Scanln(&miles)
		if miles == -1 {
			break
		}

		fmt.Print("Enter gallons used: ")
		fmt.Scanln(&gallons)

		if gallons <= 0 {
			fmt.Println("Gallons must be greater than 0. Try again.")
			continue
		}

		mpg := miles / gallons
		fmt.Printf("Miles per gallon for this trip: %.2f MPG\n", mpg)

		totalMiles += miles
		totalGallons += gallons
		tripCount++

		if totalGallons > 0 {
			fmt.Printf("Total MPG after %d trips: %.2f MPG\n", tripCount, totalMiles/totalGallons)
		}
	}

	fmt.Println("Program terminated. Have a great day!")
}
