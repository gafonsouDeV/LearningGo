# Generics

Allow functions and types to work with different data types while maintaining type safety. Enable reusable code without sacrificing performance.

Generics are introduced to solve code duplication when working with multiple types.

## Generic Functions

Functions working with multiple types using type parameters in square brackets like ``func FunctionName[T any](param T) T``. Enable reusable algorithms maintaining type safety.

## Generic Types / Interfaces

Create reusable data structures and interface definitions working with multiple types. Define with type parameters like ``types Holder[T any] struct { value T }``. Enable type-safe containers, generic slices, maps, and custom structures.

## Types Constrains

Specify which types can be used as type arguments for generics. Defined using interfaces with method signatures or type sets. Common constraints include ``any``, ``comparable`` and custom constraints

## Type Inference

Allows compiler to automatically determine generic type arguments based on function arguments or context. Reduces need for explicit type specification while maintaining type safety. Makes generic functions cleaner and more readable by eliminating redundant type specifications.

## Code example using Generics

```go
type Number Interface {
    int64 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
    var result V

    for _, v := range m {
        s += v
    }

    return s
}

func main() {
    ints := map[string]int64 {
        "first": 19,
        "second": 22
    }

    floats := ma[string]float64 {
        "first": 3.2
        "second": 4.4
    }

    Println(SumNumbers[string, int64](ints)) // 41
    Println(SumNumbers[string, float64](floats)) // 7.6
}
```
