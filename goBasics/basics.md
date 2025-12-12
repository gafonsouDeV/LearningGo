# Language Basics

## Variables & Constants

I asume that most of the readers know what is a variable and what's a constant but just in case:

- variables store in memory changeable values.
- constants store in memory unchangeable values.

In go variables can be specific typed or use type inference.(constants must be determinable at the compile-time).

### var and :=

Go provides 2 ways to declare variables.

To make explicit declaration you should use the ``var`` keyword. also it can be use inside and outside functions and if no value is provided Go assigns a default zero depending of the type of variable.

``` go
var a int = 3
var b int
```

The other way of declaring variables is using the shorthand ``:=``. It infers the type from the given value and can only be used inside functions.

``` go
a := 5 // This will infer to an integer value
```

### const and iota

Constant are declared with ``const``

```const APPLE_PRICE = 2```

The iota identifier is used in const declarations to simplify definitions of incrementing numbers. The value of iota is reset to 0 whenever the reserved word 'const' appears in the source and incremented by one after each ConstSpec.

``` go
const (
    Sunday String = iota // sunday = 0
    Monday // monday = 1
    Tuesday // tuesday = 2
)
```

### Scope and Shadowing

The scope determines the access level of the variables, it can from global to block level.

Shadowing occurs when a variable declared in a nested scope has the same name as a variable declared in an outer scope. In such cases, the variable in the inner scope shadows or hides the variable in the outer scope

``` go
func main() {
    x := 1  // Outer variable
    fmt.Println(x) // Prints: 1
    if true {
        x := 3 // Inner variable
        fmt.Println(x) // Prints 3
    }
    fmt.Println(x) // Print 1 (outer x is again accessible)
}
```

## Data types

### Numeric Types

There are 3 types of numeric types on golang.

- Integers
    -

    Signed Integers(``int8``, ``int16``, ``int32``, ``int64``)handle positive/negative numbers(1,3,-2,-343...).

    Unsigned Integers(``uint8``, ``uint16``, ``uint32``, ``uint64``) handle only non-negative but larger positive range because it doesn't need the bit of sign(2, 300, 129428492...).

- Floating Points
    -

    It can be two types of floats, ``float32``(single precision) and ``float64``(double precision). Represent real numbers using IEEE 754 standard.

- Complex numbers
    -

    Built-in support with ``complex64`` and ``complex128`` types. Create using ``complex()`` function or literals like ``3+4i``. It has it owns useful functions for mathematical computation like ``real()``, ``imag()``, ``abs()``...

### Booleans

The ``bool`` type represent ``true`` or ``false``values. Essentials for contitions, control flow and binary states. Also they're the results of logical operations like ``&&``(and), ``||``(or), ``!``(not). Or results from comparison, ``==``(equal), ``!=``(different)

### Runes

Represent Unicode code points as ``int32`` type. The definition of runes in golang could be a little confuse but making it simple is the denomination of ``char`` in other languages using UTF-8. An example could be``'⌘'`` or ``'A'``.

### Strings

- Raw String Literals
    -

    Enclosed in backticks (`) and interpret characters literally without escape sequences. Ideal for regex, file paths, SQL queries, JSON templates, and multi-line text where escaping would be extensive.

- Interpreted String Literals
    -

    Enclosed in double quotes (") and process escape sequences like \n, \t, \". Support Unicode characters and formatting. Most common string type, ideal for text needing control characters but requiring escaping of special characters.

### Type Conversion

Convert values between different types using ```Type(value)``` syntax. Go requires explicit conversion even between related types like int and ``int64``.

Example:

``` go
var n1 int = 10
var f1 float64 = float64(n1) // we transform the integer onto a real number
var uint = uint(f1) // and here we convert the real number to a unsigned integer
```
