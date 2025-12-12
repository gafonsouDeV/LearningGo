# Pointers

Variables storing memory addresses of other variables. Enable efficient memory usage and allow functions to modify values sended as parameters(In previous chapters we state that the parameters was inside the functions were copies from outside). Declared with ``*Type``, address obtained with ``&``. Pointers are essentials for performance and building data structures. Unlike other languages like C, Go has no pointer arithmetic.

```go
func main() {
    num1, num2 := 22, 100

    var pNum1 *int = &num1 // point to num1

    fmt.Println(*pNum1) // read through the pointer
    fmt.Println(pNum1) // read num1 memory address

    *pNum1 = 12
    fmt.Println(*pNum1) // print new value 12


    pNum2 := &num2 // point to num2
    *pNum2 = *pNum2 / 10
    fmt.Println(num2) // New value: 10

}
```

The ``*`` operator denotes the pointe's underlying value. This is known as "dereferencing" or "indirecting".

## Pointers with Structs

Pointers to structs enable efficient passing of large structures and allow modification of struc fields. Access fields with ``(*ptr).field`` or shorthand ``ptr.field``. Common for method receivers and when structs need to be modified by functions.

``` go
type Point struct {
    X int
    Y int
}

func main() {
    point := Point{1, 2}

    p := &v
    p.X = 10

    fmt.Println(point) // (10, 2)
}
```

## Pointers with Maps & Slices

Maps and slices are reference types - passing them to functions doesn't copy underlying data. Modifications inside functions affect original. No need for explicit pointers. Howeverm reassigning the slice/map variable itself won't affect caller unless using pointer.
