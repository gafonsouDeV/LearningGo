# Methods

Go does not have classes. However, you can define methods on types. A method is a function with a special receiver argument. The receiver appears in its own argument list between the ``func`` keyword and the method name.

```
type Point struct {
    X, Y float64
}

func (p Point) Abs() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
    p := Point{4,3}

    fmt.Println(p.Abs()) // 5
}
```

In this example, the ``Abs`` method has a receiver of type ``Point`` named p.

## Pointer Receivers

Methods receive pointer to struct rather than copy using ``func (p *Type) methodName()`` syntax. Necessary when method modifies receiver state or struct is large.

```
type Point struct {
    X, Y float64
}

func (p Point) Abs() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) Scale(num float64) {
    p.X = p.X * num
    p.Y = p.Y * num
}

func main() {
    p := Point{4,3}

    fmt.Println(p.Abs()) // 5

    p.Scale(10)
    fmt.Println(p) // (40, 30)
}
```

The ``Scale`` method must have a pointer receiver to change the ``Point`` vallue declared in the main function.
```
func (p *Point) Scale(num float64) {
    p.X = p.X * num
    p.Y = p.Y * num
}

func main() {
    p := Point{4,3}

    fmt.Println(p.Abs()) // 5

    p.Scale(10)
    fmt.Println(p) // (40, 30)
}
```

The ``Scale`` method must have a pointer receiver to change the ``Point`` vallue declared in the main function.

## Value Receivers

Methods receive a copy of struct rather than pointer. Appropiate when method doesn't modify receiver or struct is small. Can be called on both values and pointers whit Go automatically dereferencing

## Choosing a value or pointer receiver

There are two main reasons to use a pointer receiver

1. The method can modify the value that its receiver points to
2. Avoid copying the value on each method call. This can be more efficient if the receiver is a large struct.