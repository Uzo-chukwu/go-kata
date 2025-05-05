package main

import (
	"fmt"
)

func main() {

	itemValues := map[int]float64{
		1: 239.99,
		2: 129.75,
		3: 99.95,
		4: 350.89,
	}

	var totalSales float64

	fmt.Println("Enter the number of each item sold:")
	for item, price := range itemValues {
		var quantity int
		fmt.Printf("Item %d ($%.2f): ", item, price)
		fmt.Scanln(&quantity)
		totalSales += float64(quantity) * price
	}

	baseSalary := 200.0
	commission := totalSales * 0.09
	totalEarnings := baseSalary + commission

	fmt.Printf("\nTotal Sales: $%.2f\n", totalSales)
	fmt.Printf("Total Commission (9%%): $%.2f\n", commission)
	fmt.Printf("Final Earnings: $%.2f\n", totalEarnings)
}
