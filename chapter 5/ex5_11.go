package main

import "fmt"

func main() {
	n := 10

	// (a) Left-aligned triangle
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	// (b) Right-aligned inverse triangle
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	// (c) Right-aligned triangle
	for i := 1; i <= n; i++ {
		for s := 1; s <= (n - i); s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()

	// (d) Right-aligned inverse triangle
	for i := n; i >= 1; i-- {
		for s := 1; s <= (n - i); s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
