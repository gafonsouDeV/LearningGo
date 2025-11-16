package main

import "fmt"

func readInt(msg string) int {
	var number uint

	fmt.Println(msg)

	if _, err := fmt.Scanln(&number); err != nil {
		return -2
	}

	return int(number)
}

func readNumber(msg string) float64 {
	var number float64

	fmt.Println(msg)

	for {
		_, err := fmt.Scanln(&number)

		if err == nil {
			break
		}
		fmt.Println("That's not a number")
	}

	return number
}
