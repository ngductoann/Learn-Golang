# Document

<!--toc:start-->

- [Document](#document)
  - [Type constraints example](#type-constraints-example)
    - [A non-generic function](#a-non-generic-function)
    - [Add a generic function to handle multiple types](#add-a-generic-function-to-handle-multiple-types)
    - [Remove type arguments when calling the generic function](#remove-type-arguments-when-calling-the-generic-function)
    - [Declare a type constraint](#declare-a-type-constraint)
    - [Concrete type vs interface type](#concrete-type-vs-interface-type)

## Type constraints example

[Tutorial](https://go.dev/doc/tutorial/generics)

### A non-generic function

You're declaring two functions instead of one because you're working with two different types of map: one that stores int64 values, and one that stores float64 values.

```go
// SumInt64 adds together the values of m
func SumInts(m map[string]int64) int64 {
    var s int64

    for _, v := range m {
        s += v
    }

    return s
}

// SumFloat64 adds together the values of m
func SumFloats(m map[string]float64) float64 {
    var s float64

    for _, v := range m {
        s += v
    }

    return s
}
```

In this code, you:

- SumFloats takes a map of string to float64 values.
- SumInts takes a map of string to int64 values.

At the top of main.go, beneath the package declaration, paste the following main function to initialize the two maps and use them as arguments when calling the functions you declared in the preceding step:

```go
func main() {
        // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))
}
```

### Add a generic function to handle multiple types

You'll add a single generic function that can receive a map containing either integer or float values, effectively replacing the two functions you just wrote with a function.

```go
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V

    for _, v := range m {
        s += v
    }

    return s
}
```

In this code, you:

- Declare a SumIntsOrFloats function with two type parameters (inside the square brackets), K and V, and one arguments that use the type parameters, m of type `map[K]V`. The function returns a value of type V.
- Specify for the K type parameter the type constraint comparable. Intended specifically for cases like these, the comparable constraint is predeclared in Go. It allows any type whose values may be used as an operand of the comparison operators == and !=. Go requires that map keys be comparable. So declaring K as comparable is necessary so you can use K as the key in the map variable. It also ensures that calling code uses an allowable type for map keys.
- Specify for the V type parameter a constraint that is a union of two types: int64 and float64. Using | specifies a union of the two types, meaning that this constraint allows either type. Either type will be permitted by the compiler as an argument in the call code.
- Specify that the m argument is of type `map[K]V`, where K and V are the types already specified for the type parameters. Note that we know `map[K]V` is a valid map type because K is a comparable type. If we hadn't declared K comparable, the compiler would reject the reference to `map[K]V`.

In main.go, beneath the code you already have, paste the following code.

```go
fmt.Printf("Generic Sums: %v and %v\n",
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))
```

In this code, you:

- Call the generic function you just declared, passing each of the maps you created.
- Specify type arguments - the type names in square brackets - to be clear about the types that should replace type parameters in the function you're calling.
- Print the sums returned by the function.

### Remove type arguments when calling the generic function

You can omit type arguments in calling code when Go compiler can infer the types you want to use. The compiler infer type arguments from the types of function arguments.

Note that this isn't always possible. For example, if you needed to call a generic function that had no arguments, you would need to include the type arguments in the function call.

**Write the code:**

- In main.go, beneath the code you already have, paste the following code:

```go
fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
    SumIntsOrFloats(ints),
    SumIntsOrFloats(floats))
```

In this code, you:

- Call the generic function, omitting the type arguments.

### Declare a type constraint

You declare a type constraint as an interface. The constraint allows any type implementing the interface. For example, if you declare a type constrain interface with these three methods, then use it with a type parameter in a generic function, type arguments used to call the function must have all of those methods.

Constrain interfaces can also refer to specific types, as you'll se in the section.

**Write the code**

1. Just above main, immediately after the import statement, paste the following code to declare a type constraint.

```go
type Number interface {
    int64 | float64
}
```

In this code, you:

- Declare the Number interface type to use as a type constraint.
- Declare a union of int64 or float64 inside the interface.

Essentially, you're moving the union from the function declaration into a new type constraint. That way, when you want to constrain a type parameter to either int64 or float64, you can use this Number type constraint instead of writing out int64 | float64.

### Concrete type vs interface type

In Go programming language, a **"concrete type"** is **a type that you can directly instanctiate or create a value from.** This means that `the type can directly represent a set of values and that you can create an instance of this type without any additional information.`
**`Example of concrete types`** in Go include `int`, `bool`, `string`, `float64`, arrays, slices, maps, and structs, among others. These types have specific, predetermined storage layouts and characteristics.
This is in constrat to an **"interface type"** in Go, which defines a contract ( set of methods or types) but does not represent a specific data layout or instance. `Interface types are abstract - they represent behavior or type but not a specific set of values`

**`Concrete type`**

For example, you can declare a struct with a specific set of fields and then create an instance of this struct. `This is a **concrete type**`

```go
type Employee struct {
  Name string
  Age int
}

emp := Employee{
    Name: "John Doe",
    Age: 30,
}
```

**`Interface type`**

But if you declare an interface like `io.Reader`, you cannot directly create an instance of `io.Reader`. Instead, you create instances of concrete types that satisfy the `io.Reader` interface.

```go
type Reader interface {
  Reader(p []byte)(n int , err error)
}
```

In the latter case, `io.Reader` is an **interface type**, and any type that implements the `Read` method with the correct signature is said to "satisfy" the `io.Reader` interface.
