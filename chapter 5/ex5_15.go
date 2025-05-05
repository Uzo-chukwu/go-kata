package main

import "fmt"

func main() {
	n := 5

	// Upper part of the shape
	for i := 1; i <= n; i++ {
		for j := 1; j <= (2*i - 1); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Lower part of the shape
	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= (2*i - 1); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
