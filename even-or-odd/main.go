package main

import "fmt"

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range values {
		if value%2 == 0 {
			fmt.Printf("Even %v\n", value)
		} else {
			fmt.Printf("Odd %v\n", value)
		}
	}
}
