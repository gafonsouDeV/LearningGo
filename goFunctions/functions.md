# Functions

Declared with ``func`` keyword, support parameters and return values. Can be assigned to variables and support multiple return values, named returns, and variadic parameters.

A function can take zero or more arguments

```
package main

import "fmt"

func add(num1 int, num2 int) int {
    return num1 + num2
}

func main() {
    fmt.Println(add(41+22)) // 63
}
```

The functions are fundamental building blocks for organizing code logic.

## Variadic Functions

Functions accepting variable number of arguments of samy type.

``func name(args ...Type)``

Arguments treated as slice inside function.

The ``...`` can be used also to unpack slices, arrays...

```
package main

import "fmt"

func sum(nums ...int) int {
    var result int
    for _, num := range nums {
        result += num
    }

    return result
}

func main() {
    nums := [3]int {2,5,3}
    fmt.Println(sum(nums...)) // 10
}
```


## Multiple return Values

Go functions can return multiple values, usually used for returning result and error.

``func name() (Type1, Type2) {}``

Caller receives all returned values or uses blank identifier to ignore unwanted values.

This is the idiomatic for error handling pattern

Here two examples of how to use it:

```
func calculateSumMaxAvg(numbers []int) (int, int, float64) {
    numbSize = len(numbers)
    if numbSize == 0 {
        return 0, 0, 0.0
    }

    sum := 0
    max := numbers[0]

    for _, num := range numbers {
        sum += num

        if num > max {
            max = num
        }
    }

    avg := float64(sum) / float64(numbSize)
 }

 func main() {
    numbers := [5] int{ 2, 4, 3, 5, 1}

    totalSum ,maxNumb, avg = calculateSumMaxAvg(numbers)

    fmt.Println(totalSum, maxNumb, avg) // 15, 5, 3
 }
```

And now an example with error handle

```
func divide(num, num2 float32) (float32, error) {
    if num2 == 0 {
        return 0, errors.New("Division by zero is not possible")
    }

    return a / b, nil
}
```


## Anonymous Functions

Functions declared withoud names as in some languages it can be also named as lambdas. Can be assigned to variables, passed as arguments or executed immediately. Useful for short operations, callbacks, goroutines and closures. Access enclosing scope variables. Is it commonly use in event handlers and functional patterns.


``` 
func main() {
    sum := func(num1 int, num2 int) {
        return num1 + num2
    }

    fmt.Println(sum(2,5)) // 7
}
```

## Closures

Functions capturing variables from sorrounding scope, accessible even after outer function returns. Useful for event handling, iterators, functional programming...

To be more specific a closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables.

'''
func adder() func(int) int {
    sum := 0
    return func (num int) int {
        sum += num
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()

    for i := 0; i < 5; i++ {
        fmt.Println(pos(i),neg(-i))
    }
}
'''

The output will be:

0 0 // 0 + 0

1 -1 // 0 + 1

3 -3 // 1 + 2

6 -6 // 3 + 3

10 -10 // 6 + 4

Each closure is bound to its own sum variable


## Named Return Values

Function return parameters can be named and treated as variables within function. The ``return`` statement without arguments returns current values of named parameters.

func f() (num int, word string) {
    num := 10
    word := "bazanga"
    return // same as return num, word
}

## Call by Value

Go creates copies of values when passing to functions, not references to originals, this will be explain further on this learning guide. It provides safety but can be expensive for large data.