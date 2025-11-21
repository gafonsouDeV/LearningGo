package utils

import (
	"bufio"
	"fmt"
	"os"
)

type Number interface {
	int32 | float32
}

func ReadNumber[T Number](msg string) (T, error) {
	var number T
	fmt.Println(msg)

	if _, err := fmt.Scanln(&number); err != nil {
		return -2, fmt.Errorf("error to read %T", number)
	}

	return number, nil
}

func ReadString(msg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	text, _ := reader.ReadString('\n')
	return text
}

func ReadValidatedNumber[T Number](msg string) T {
	var value T
	for {
		v, err := ReadNumber[T](msg)
		if err == nil {
			value = v
			break
		}
		fmt.Println("Invalid input, try again")
	}
	return value
}
