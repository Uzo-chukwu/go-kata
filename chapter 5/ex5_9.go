package main

import "fmt"

func main() {
	var sum int64 = 0

	fmt.Println("n\tSum")
	fmt.Println("----------------")

	for n := int64(1); n <= 100; n++ {
		sum += n
		fmt.Printf("%d\t%d\n", n, sum)
	}
}
