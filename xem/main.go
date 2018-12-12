package main

import "fmt"

func main() {
	var xem float32 = 10000
	var balance float32
	fmt.Println(xem, balance)

	var i = 0
	for {
		balance = equation(i, xem)
		fmt.Printf("Day %c: %f", i, balance)

		if balance > 10000 {
			break
		}

		i++
	}
}

func equation(i int, b float32) float32 {
	if i == 1 {
		return b * 0.1
	}

	minusOne := equation(i-1, b)

	return minusOne + (b-minusOne)*0.1
}
