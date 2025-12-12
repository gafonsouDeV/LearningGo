# Data Structures

## Arrays

Fixed-size sequences of same-type elements.

``[n]T`` is an array of n values of type T.

If you declare the array only with the lenght it is initialized to zero values in all the array positions. ``var nums [10]int``. Or you can declare it completly:  ``week := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}``.

## Slices

Dynamic arrays built on top of arrays. Reference types with length and capacity.

### Capacity and Growth

Slice capacity determines when reallocation occurs during append operations. Go usually doubles capacity for smaller slices. Pre-allocate with ``make([]T, length, capacity)`` to optimize memory usage and minimize allocation in performance-critical code.

``` go
s := make([]int, 2, 3)
s = append(s, 0)
ftm.Print("Capacity: ", cap(s)) // Capacity: 3
s = append(s, 0)
ftm.Print("Capacity: ", cap(s)) // Capacity: 6
```

### make()

Is a built-in function that creates and initializes slices, maps, and channels. Unlike ``new()`` returns usable values.

The make function allocates a zeroed array and returns a slice that refers to that array.


### Slice to Array Conversion

Convert slice to array using ``[N]T(slice)`` let's take the above example:

``` go
weekArr := [7]string(week)
```
This copies data from slice to fixed-size array. Panics if slice has fewer than N elements.

### Array to Slice Conversion

Convert array to slices using expresions like ``array[:]`` or ``array[start:end]``. Creates slice header pointing to array memory, no data copying.

``` go
arr := [5]int{1,2,3,4,5}
slice := arr[1:4] // Creates a slide from index 1 to 3
```

Modifications through slice affect original array.

## Strings

Immutable sequences of bytes representing UTF-8 encoded text. String operations create new strings instead of modifying existing ones. Iterate by byte(indexing) or runes(for range).

## Maps

Built-in associative data mapping keys to values. Keys must be comparable types.

``make(map[KeyType]ValueType)``

Map types are reference types, like slices and when initialize the default value is nil. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic;

``` go
var week map[int]string // nil default value

week = make(map[int]string) // map initialized

week[1] = "Monday" //assign Monday to key 1

l := len(week) // the build in len function returns the len of the map

delete(week, 1) // the build in delete function removes an entry from the map 
```

### Comma-Ok Idiom

Is a pattern for safely testing map key existence or type assertion success using ``value, ok := map[key]`` or ``value, ok := Interface(Type)``. Returns both value and boolean status, preventing panics and distinguishing zero values from missing keys.

``` go
msg := "This is a message"
n, ok := msg.(int) 
if ok {
    fmt.Println("msg: ", msg)
} else {
    fmt.Println("not an int") // In this case it will show this 
}

var stocks map[string]float64
sym := "TTWO"
price := stocks[sym]
fmt.Printf("1. %s -> $%.2f\n", sym, price)

// using comma ok to check if key exists
if price, ok := stocks[sym]; ok {    
    fmt.Printf("2. %s -> $%.2f\n", sym, price)
} else {
    fmt.Printf("2. %s not found\n", sym)
}

```

## Structs

Custom data types related fields under single name. Similar to classes but methods defined separately. It is used to create complex data models, organize information, define application data structure. 

The oficial definition in golang "A ``struct`` is a collection of fields"

``` go
type Point struct {
    x int
    y int
}

func main() {
    fmt.Println(Point{1,2})
}
```

### Struct Tags & JSON

Structs tags provides metadata using backticks with key-value pair. JSON tags control field names. You can check the list of [Struct Tags](https://go.dev/wiki/Well-known-struct-tags).

This is Essential for APIs and data serialization formats.

``` go
type User struct {
    name string `json:"name"`
    email string `json:"email"`
}
```

### Embedding Structs

Struct embedding includes one struc inside another without field names, making embedded fields directly accessible. It has a similar behavior with the inheritance, but is not the same. Provides composition-based design.

``` go
type User struct {
    name string
    email string
}

type Admin struct {
    User
    adminCode uint32
}
```
