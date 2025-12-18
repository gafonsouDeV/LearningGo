package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Is even number")
		} else {
			fmt.Println(i, " Is odd number")
		}
	}
}
