package utils

import "fmt"

func ValidateNumberInput(input int) bool {
	return input <= 0
}

func ReadNumberInput() int {
	var input int

	for {
		_, err := fmt.Scan(&input)
		if err == nil && ValidateNumberInput(input) {
			break
		}
		fmt.Println("Please enter a valid positive number:")
	}

	return input
}
