package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
