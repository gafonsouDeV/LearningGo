package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func menu() {
	operation := -1
	var result float64 = 0

	for {
		showMenuText(result)
		operation = readInt("Enter the operation")

		if operation == 0 {
			break
		}

		if operation == -2 {
			result = 0
		}

		result = calculateOperationSelected(uint(operation), result)

		cleanMenu()
	}

	fmt.Println("Thanks for using this calculator")
}

func showMenuText(result float64) {
	fmt.Printf("Calculator\nresult: %v\n", result)
	fmt.Println("Enter the number of the Operations:\n1: Add\n2: Substract\n3: Division\n4: Multiplication")
	fmt.Println("Not a number to reset\n0: Exit")
}

func cleanMenu() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // Linux, macOS
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
