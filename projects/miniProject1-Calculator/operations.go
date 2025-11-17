package main

import "fmt"

func calculateOperationSelected(operation uint, result float64) float64 {
	// anonymous function for simple operations
	substract := func(num1 float64, num2 float64) float64 { return num1 - num2 }
	division := func(num1 float64, num2 float64) float64 { return num1 / num2 }
	multiplication := func(num1 float64, num2 float64) float64 { return num1 * num2 }

	// map of the operations
	operations := map[uint]func(float64, float64) float64{
		1: add,
		2: substract,
		3: division,
		4: multiplication,
	}

	if option, exist := operations[operation]; exist {
		num1 := result
		var num2 float64
		if result == 0 {
			num1 = readNumber("Enter the first number")
		}
		num2 = readNumber("Enter the second number")

		return option(num1, num2)
	} else {
		fmt.Println("Invalid option")
	}

	return result
}

// addition using a normal function
func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}
