# Interfaces

As in other languages define contracts specifying method signatures without implementation. Types satisfy interfaces implicitly by implementing required methods. Enable polymorphism and loose coupling. Empty interface ``interface {}`` accepts any type.

## Interface Basics

Declared with ``type InterfaceName interface{}`` sintax. Enable polymorphism and flexible, testable code depending on behavior rather than concrete types.

```
type Geometry interface {
    area() float64
    perim() float64
}

type Rectangle struct {
    width, height float64
}

type Circle struct {
    radius float64
}

func (rect Rectangle) area() float64 {
    return rect.width * rect.height
}

func (rect Rectangle) perim() float64 {
    return 2*rect.with + 2*rect.height
}

func (circle Circle) area() float64 {
    return math.Pi * circle.radius * circle.radius
}

func (circle Circle) perim() float64 {
    return 2*math.Pi*circle.radius
}

func measure(g Geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Printlne(g.perim())
}

func main() {
    r := Rectangle{with:2, height:5}
    c := Circle{radius: 3}

    measure(r)
    measure(c)
}
```

## Empty Interface

The empty interface can hold values of any type since every type implements at least zero methods. Requires type assertions or type swtiches to access underlying values. Common in APIs handling unknown data types.

```
func describe(i interface{}) {
    fmt.Println(i)
}

func main() {
    var i interface{}

    describe(i)

    i = 10
    describe(i)

    i = "Hello"
    describe(i)
}
```

## Embedding Interfaces

Create new interfaces by combining existing ones, promoting composition and reusability. Embedded interface methods automatically include. Enables interface hierarchies. Supports composition over inheritance.

## Type Assertions

Extract underlying concrete value from interface ``value.(Type)`` or ``value, ok := value.(Type)`` for save assertion. Panics if type assertion fails without ok form. 

```
func main() {
    var i interface{} = "hello"

    s := i.(string)
    fmt.Println(s)

    s, ok := i.(string)
    fmt.Println(s, ok)

    num, ok := i.(int32)
    fmt.Println(num, ok)

    num = i.(int32) // panic
    fmt.Println(num)
}
```

If the interface holds the type then ok will be true and the varible the underlying value. If not, ok will be false, variable will be the zero value of type ``T``, and no panic occurs.

## Type Switch

Special form of switch statement that operates on types rather than values ``switch val := i.(type)``. Used with interfaces to determine underlying concrete rype. Each case specifies types to match. 

```
func main() {
    do(10)
    do("Dumbo")
    do(true)
}

func do(i interface{}) {
    switch v := i.(type) {
        case int:
            fmt.Println(v * 2)
        case string:
            fmt.Println(len(v))
        default:
            fmt.Println(v)
    }
}
```