package main

import "fmt"

func main() {
	var name string
	var grade string
	grades := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0}

	fmt.Println("Enter student names and grades (A, B, C, or D):")

	for i := 1; i <= 5; i++ {
		fmt.Printf("Student %d name: ", i)
		fmt.Scanln(&name)
		fmt.Printf("Student %d grade: ", i)
		fmt.Scanln(&grade)

		switch grade {
		case "A", "B", "C", "D":
			grades[grade]++
		default:
			fmt.Println("Invalid grade! Please enter A, B, C, or D.")
			i-- // Retry input
		}
	}

	fmt.Println("\nGrade Summary:")
	fmt.Printf("A: %d students\n", grades["A"])
	fmt.Printf("B: %d students\n", grades["B"])
	fmt.Printf("C: %d students\n", grades["C"])
	fmt.Printf("D: %d students\n", grades["D"])
}
