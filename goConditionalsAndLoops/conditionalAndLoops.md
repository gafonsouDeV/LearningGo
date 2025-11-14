# Conditionals and Loops

## Conditionals

Control program flow baesd on conditions.

### if

Basic conditional statement for executing code based on boolean conditions

``` 
func main() {
    age := 10

    if age <= 18 {
        fmt.Println("You're under age")
    }
}
```

The ``else`` clause provides alternative executions when the condition is false.

``` 
func main() {
    age := 10

    if age <= 18 {
        fmt.Println("You're under age")
    } else {
        fmt.Println("You're an adult")
    }
}
```

Multiple conditions can be checked using ``else if``  clause.

func main() {
    age := 10

    if age <= 18 {
        fmt.Println("You're under age")
    } else if age >= 60 {
        fmt.Println("You're an elder")
    } else {
        fmt.Println("You're an adult")
    }
}

### switch

The switch clause is a clean way to compare variable against multiple values and execute corresponding code block. In go no break statements needed.

```
var weekDay string

switch day {
    case 1:
      weekDay = "Monday"
    case 2:
      weekDay = "Tuesday"
    case 3:
      weekDay = "Wednesday"
    case 4:
      weekDay = "Thursday"
    case 5:
      weekDay = "Friday"
    case 6:
      weekDay = "Saturday"
    case 7:
      weekDay = "Sunday"
    default:
      panic("That's not a day of the week")
}
```

## Loops

Go has only one looping construct, the ``for`` loop. 

The basic ``for`` loop has three components separated by semicolons. The init statement, the condition and the post statement

```
func main() {
    sum := 0
    for i := 0; i < 3; i++ {
        sum += 1
    }

    fmt.Println(sum) // 3
}
```

### for-range

This is a special form of loop for iterating arrays, maps, slices, strings, and channels. The blank identifier ``_`` is used to ignore unwanted return values.

To use this kind of ``for`` loop, the ``range`` clause is used

```
week := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

for i, value := range week {
    fmt.Printf("array value at [",i,"]=",value)
}

weekMap := map[string]int{"Monday": 0, "Tuesday": 1, "Wednesday": 2, "Thursday": 3, "Friday": 4, "Saturday": 5, "Sunday": 6}

for key, value := range week {
    fmt.Printf(key,": ",value)
}

// to ignore the key or value, use the blank _ identifier

for _, value := range week {
    fmt.Printf(value)
}
```

To iterate the strings first it is important to understand how they work in golang. Unlike some programming languages, Go treats strings as immutable values, which means once a string is created it cannot be modified. The string is a sequence of Unicode characters represented as a read-only slice of bytes and to proper character handling it should be use the ``rune`` type


```
text := "String to iterate"

for index, runeValue := range text {
    fmt.Println("Index: %d, Character: $c", index, runeValue)
}
```

### break & continue

The ``break`` clause exits inmmediately innermost loop or switch statement. If there's nested loops it will exit the inmediate loop unless used with labels to break outer loops.

```
package main

import "fmt"

func main() {
	for outer := 0; outer < 5; outer++ {
		if outer == 3 {
			fmt.Println("Breaking out of outer loop")
			break 
		}
		fmt.Println("The value of outer is", outer)
		for inner := 0; inner < 5; inner++ {
			if inner == 2 {
				fmt.Println("Breaking out of inner loop")
				break 
			}
			fmt.Println("The value of inner is", inner)
		}
	}
}
```

The output of that snippet of code:

The value of outer is 0

The value of inner is 0

The value of inner is 1

Breaking out of inner loop

The value of outer is 1

The value of inner is 0

The value of inner is 1

Breaking out of inner loop

The value of outer is 2

The value of inner is 0

The value of inner is 1

Breaking out of inner loop

Breaking out of outer loop

The ``continue`` statement is used when you want to skip the remaining portion of the loop and went to the next iteration.

```
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Continuing loop")
			continue // break here
		}
		fmt.Println("The value of i is", i)
	}
}
```

In this case the output will be:

The value of i is 0

The value of i is 1

The value of i is 2

The value of i is 3

The value of i is 4

Continuing loop

The value of i is 6

The value of i is 7

The value of i is 8

The value of i is 9
