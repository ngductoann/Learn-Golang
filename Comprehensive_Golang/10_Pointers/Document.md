# Document

## What are pointers ?

All values are stored in memory:

- post office boxes are a good analogy for memory
  - Every location in memory has an address
  - **A pointers is a memory address**
  - "Here's the value AND what is the address ?"
- lexical element:

  - `&`
  - `*`
  - `*int`

In the Go programming language, **a pointer** refers to **a variable that holds the memory address.** Pointers are a powerful feature that allows you to **directly manipulate memory** and build complex data sturctures like linked lists, trees, and more. However, they also require careful use, as **improper pointer usage can lead to bugs and errors.**

There are two fundamental operations involving pointers:

1. "`Address operation(&)`:" The & operator is used to get the memory address of a variable. If `x` is an integer variable then `&x` will give you a pointer to `x` -- that is, a memory address where the integer `x` is stored.

2. **`Dereferencing operator(*)`:** The * operator is used to get the value stored at a memory address. If `p` is a pointer to an integer then `*p`gives you the integer that`p` points to.

```go
func main() {
    i := 42
    p := &i // p is a pointer to i
    fmt.Println(*p) // Dereferencing 'p' gives you the integer that 'p' points to
    *p = 21 // You can change the value that 'p' points to
    fmt.Println(i) // Now i is 21
}
```

In this example, `p` is a pointer to `i`. You can get the value of `i` by derefercening `p` with `*p`, and you can change the value of `i` by assinging to `*p`.

Pointers also come into play when you're dealing with functions. If you pass a variable to a function in Go, the function gets a copy the variable -- *everything in go is pass by value* -- if the function changes the variable, it won't affect the original. But if you pass a pointer to a variable, the function can dereference the pointer to change the original variable.

Lastly, pointers are curcil when dealing with structs in Go. Since Go doesn't have classes and objects, you can use structs to create complex types, and you can use pointers to structs to modify these structs or to avoid copying the structs around.

## Seesing type & value for pointers

You can print the type and value of a pointer in Go by using the `fmt.Printf` function along with the `%T` and `%v` format verbs. The `%T` verb is used to print the type of a variable, while the `%v` verb to print the value in its default format.

For example, if you have a pointer to an integer, you can print its type and value like this:

```go
package main

import "fmt"

func main() {
    num := 42
    numPtr := &num
    fmt.Printf("Type of numPtr: %T\n", numPtr)
    fmt.Printf("Value of numPtr: %v\n", numPtr)
}
```

In this code, `numPtr` is a pointer to an integer. The `%T` verb will print the type of `numPtr`, which is `*int`, and the `%v` verb will print the value of `numPtr`, which is the memory address of `num`.

## Dereferencing pointers

Continuing the description from above, if you want to print the value that the pointer is pointing to, you would need to **dereference the pointer using `*`** like this:

```go
fmt.Printf("Value pointed to by numPtr: %v\n", *numPtr)
```
