package main

import "fmt"

func main() {
	n := 1

	switch n {
	case 1:
		fmt.Println("The number is 1")
	case 2:
		fmt.Println("The number is 2")
	default:
		fmt.Println("The number is not 1 or 2")
	}
}
