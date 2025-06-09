# Document

<!--toc:start-->

- [Document](#document)
  - [What are pointers ?](#what-are-pointers)
  - [Seesing type & value for pointers](#seesing-type-value-for-pointers)
  - [Dereferencing pointers](#dereferencing-pointers)
  - [Pass by value, pointers / reference types, and mutability](#pass-by-value-pointers-reference-types-and-mutability)
  - [Pointer & value semantics defined](#pointer-value-semantics-defined)
  - [Pointer & value semantics heuristics](#pointer-value-semantics-heuristics)
  - [Pointers, values, the stacks, & the heap](#pointers-values-the-stacks-the-heap)
  - [Exploring method sets path 1](#exploring-method-sets-path-1)
  - [Exploring method sets path 2](#exploring-method-sets-path-2)
  <!--toc:end-->

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

In the Go programming language, **a pointer** refers to **a variable that holds the memory address.** Pointers are a powerful feature that allows you to **directly manipulate memory** and build complex data structures like linked lists, trees, and more. However, they also require careful use, as **improper pointer usage can lead to bugs and errors.**

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

In this example, `p` is a pointer to `i`. You can get the value of `i` by derefercening `p` with `*p`, and you can change the value of `i` by assigning to `*p`.

Pointers also come into play when you're dealing with functions. If you pass a variable to a function in Go, the function gets a copy the variable -- _everything in go is pass by value_ -- if the function changes the variable, it won't affect the original. But if you pass a pointer to a variable, the function can dereference the pointer to change the original variable.

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

## Pass by value, pointers / reference types, and mutability

In the go, a **reference type** is pointer - a type of data that refers to, or "point to", the location in memory of a value. Go has several reference types, including **pointers, slices, maps, channels, functions, and interfaces.**

Here's a brief description of each:

1. **`Pointer`:** A pointer holds the memory address of a value. It allows you to directly access add modify the memory location of a value.
2. **`Slices`:** A slice is a description of an array segment. It includes a pointer to the array, the length of the segment, and its capacity (the maximum length of the segment).
3. **`Maps`:** A map is a powerful data structure that associates values of one type (the key) with values of another type (the value). It's an unordered collection of key-value pairs.
4. **`Channels`:** Channels are used for communitation between goroutine to another.
5. **`Functions`:** In Go, functions are first-class citizens, meaning they can be assigned to variables, passed as arguments to other functions, and returned as values from other functions. When a function is assigned to a variable, the variable stores a reference to the function.
6. **`Interfaces`:** An interface type represents a set of method signatures. It provides a way to specify the behavior an objects: If something can do this, then it can be used here.

Remember that **when you assign a reference type to another variable, the new variable references the same memory location.**

In Go, all data is **passed by value**, which means that whenever you pass data to a function, Go creates a copy of that data and assigns the copy to a parameter variable. The function can do whatever it wants to the copy without affecting the original data.

**Mutability = changeable**
However, **mutability** plays a critical role when you consider how "pass by value" works in combination with more **composite / aggregate / complex data** types like slices and maps, and pointers.

A **mutable** value is a value that can be changed. In Go, **slices, maps,** and **pointers** are mutable data types. Even though they are passed by value, they still behave as if they were passed by reference because **the "value" that is copied and passed is the reference to the underlying data,** not the actual data.

For example, let's say you have a slice:

```go
nums := []int{1, 2, 3}
```

And you pass it to a function that modifies it:

```go
func modify(s []int) {
  s[0] = 100
}
```

If you can this function and then print out `nums`:

```go
modify(nums)
fmt.Println(nums)
```

You'll se `[100, 2, 3]`, not `[1, 2, 3]`. This is because the slice header, which includes the pointer to the underlying way, is **passed by value**. But this copied value (the pointer) still pointer to the underlying array. So changes made inside the function are visible outside of it, even though the data was passed by value.

In the case of **pointers**, the pointer itself is **passed by value** (meaning the function gets a copy of the address), but the data it points to is the same. Therefore, **dereferencing the pointer and modifying the value it points to inside the function will modify the original value.**

## Pointer & value semantics defined

**Value Semantics**

Value semantics in Go refers to when actual data of a variable is passed to a function or assigned to another variable. This means that the new variable or function parameter gets a completely independent **copy of the data - EVERYTHING IN GO IS 'PASS BY VALUE'**.

Here's a simple example of value semantics:

```go
package main

import "fmt"

func addOne(v int) int {
  v = v + 1
  return v
}

func main() {
  i := 1
  fmt.Println(addOne(i)) // Prints: 2
  fmt.Println(i) // Prints: 1
}
```

In this example, the function `addOne` doesn't modify the original `i` variable because it's operating on a copy of `i`.

**Pointer Semantics**

Pointer semantics in Go, on the other hand, involve **passing the memory address (a "pointer") rather than the data itself.** This means that you can modify the original data, not just a copy of it. **EVERYTHING IN GO IS 'PASS BY VALUE'.**

```go
package main

import "fmt"

func addOne(v *int) {
  *v = *v =1
}

func main() {
  i := 1
  addOne(&i)
  fmt.Println(i) // Prints: 2
}
```

In this example, the `addOne` function modifies the original `i` variable, because it's operating on a pointer to `i`, not a copy, not a copy.

The choice between **pointer and value semantics**in Go depends on the specific needs of your program:

- **Value semantics** and **simpler** and often **safer** because the avoid side effects, but they can be **inefficient** for large data structures because they involve copying data.
- **Pointer semantics** can be more efficient and **flexible**, but they also require careful management to avoid **bugs** related to shared, **mutable state.**

| Value semantics                                    | Pointer semantics                                   |
| -------------------------------------------------- | --------------------------------------------------- |
| **Copied** value (copied values - "pass by value") | **Shared** memory (copied values - "pass by value") |

## Pointer & value semantics heuristics

There are some general guidelines you can follow when deciding whether to use pointer or value semantics in Go:

1. **Use Value Semantics When Possible**:

- Value semantics are **simpler** and usually **safe**, since they don't involve shared state or require you to think about memory management.
- rule of thumb, if a function doesn't need to modify int inputs, or the data you're working with in small (like built-in types or small struct), use value semantics.
- rule of thumb: use value semantics for builtin types (numeric, string, bool) and reference types (slices, maps, channels, pointer, interfaces, functions)

2. **Use Pointer Semantics For Large Data:**

- Copying large structs or arrays can be inefficient.
- If the data you're working with is large, you might want to use pointer semantics to avoid the cost of copying the data. A rule of thumb: 64 bytes or larger, use pointers.

3. Use Pointer Semantics for Mutability:

- If a function or method needs to modify its receiver or an input parameter, you'll need to use pointer semantics.
- This is a common use case for methods that need to update the state of a struct.

4. Consistency:

- It's important to be consistent. If some functions on a type use pointer semantics and others use value semantics, it can lead to confusion. Typically, once a type has a method with pointer semantics, all methods on that type should have pointer semantics.

5. Pointer Semantics When Interfacing With Other Code:

- If you're interfacing with other code (like a library or a system call), you might need to use pointer semantics. For example, the `json.Unmarshal` function in the Go standard library requires a pointer to a value to populate it with umrshalled data.

Value semantics:

- **Value semantics facilitate higher levels of integerity**
  - "The majority of bugs that we get in software have to do with the mutation of memory" ~ Bill Kennedy
  - Everybody gets their own copy of the data (pass by value) helps keep data bug free.
    - If everybody is isolated to their own copy of the data (pass by value) then the bug is isolated to that one piece of code.
- **Value semantics also reduce side-effects around concurrency**
- more likely to **keep values on the stack**
- you can have methods that use pointers if there's some form of **unmarhaling** going on.

POINTER semantics:

- you don't want to pass around a lof of data
- you want to change the data at a location (mutability)

**Everything in Go is pass by value.** Drop any pharses and concepts you may have from other languages. Pass by reference, pass by copy - forget those pharses. "Pass by value." That is the only pharses you need to know and remember. That is the only pharse you should use. Pass by value. Everything in Go is pass by value. In Go, what you see is what you get (wysiwyg). Look at what is happening. That is what you get.

## Pointers, values, the stacks, & the heap

When you use the pointer semantics, it's more likely for values to escape to the heap.

**Value semantics & the stack**
In Go, when a function receives a value (not a pointer), it gets its own copy of that value. This copy is typically placed on the **stack**, which is fast and doesn't involve any form of garbage collection. Once the function returns, this memory can be instantly reclaimed.

**Pointer semantics & the heap**
One the other hand, when a function receives a pointer or returns a pointer to a local variable, it indicates to the compiler that this value could be shared across goroutine boundaries or could persist after the function returns. To ensure that the data will remain variable, the Go compiler must allocate it on the **heap**, rather than on the stack. **Heap allocation is more expensive and requires garbase collection.**

**Escape analysis**
It's important to note that the Go's **escape analysis** decides whether a variable should be allocated on the stack or the heap. Escape analysis examines the function too see if the pointer to a variable is being returned or if the variable is captured within a function literal (closure). If variable does not escape, it may be kept on the stack.

To see escape analysis in Go, you use the `-gcflags` flag followed by `-m` when running the `go build` or `go run` command. The `-m` option instructs the compiler to print escape analysis information.

Here's an example:

```bash
go run -gcflags '-m' main.go
```

You can use `-m=2` for more detailed information:

```bash
go run -gcflags '-m=2' main.go
```

This will print out escape analysis and **inlining decisions.**

**Inlining Decisions**
Inlining in programming is an optimization that involves replacing a function call site with the body of the called function. In the Go compiler, inlining decisions refer to the choices made by the compiler about when and where to apply this optimization.

Inlining can often make a program run faster due to several reasons:

1. It can **save the overhead of a function call** (i.e, the time it takes to jumps to the function code, set up the function's local variable, and then jump back when the function is complete).
2. It can make **further optimizations** possible. Once the code from the function is inserted into the calling function, the compiler may be able to see opportunities to optimize this large block of code that weren't visible when the code was split across multiple functions.

On the other hand, inlining too much code can lead to a large binary, and possibly, due to cache effects, slower code. Hence, the Go compiler makes inlining decisions to strike a balance, inlining some small and frequently called functions while leaving large and less frequently called ones alone.

When you run the Go compiler with the `-m` flag, as in `-gcflags='-m'`, it will print information about its inlining decisions along with escape analysis information. You'll see lines like "inlining call to X" where X is a function that the compiler decided to inline.

Do remember, this information is mostly useful for understanding performance characteristics of your Go code, paritcularly around memory management (specifically, whether certain data stays on the stack or escapes to the heap).

Remember, **the decisions of using pointer or value semantics should be more about sharing and mutating state rather than memory allocation and performance**, except for performance-critical parts of the code or handling very large data structures.

## Exploring method sets path 1

In Go, a 'method set' is **the set of methods attached to a type. This concept is key to the Go's interface mechanism,** and it is associated with both the value types and pointer types.

- **The method set of a `type T` consists of all methods with `receiver type T`.**
  - These methods can be called using variable of type T.
- **The method set of a `type *T`** consists of all methods with `receiver type *T or T`.
  - These methods can be called using variables of type `*T`.
  - It can call methods of the corresponding non-pointer type as well.

The idea of the method set is **`integral to how interface are implemented`** and used in Go. An interface in Go defines a method set, any type whose method set is a superset of the interface's method set is considered to implement that interface.

A curcial thing to to remember is that in Go, **if you define a method with a pointer receiver, the method is only in the method set of the pointer type.** This is important in the context of **interfaces** because **if an interface require a method that's defined on the pointer (not the value), then you can only use a pointer to that type to satisfy the interface, not a value of the type.**

Here's an example to illustrate this:

```go
type T struct {
}

func (t T) M1() {
    // ...
}

func (t *T) M2() {
    // ...
}

type X interface {
    M1() // This method can be called on a value of type T
    M2() // This method can only be called on a pointer to T
}

func main(){
    var n X

    t := T{}
    n = t // This will be a compiler error, because T does not implement M2()

    tPtr := &T{}
    n = tPtr // This is fine, because *T implements both M1() and M2()
}
```

In this example, `T` implements `M1()` and `*T` implements `M1()` and `M2()`. Therefore, `*T` satisfies interface `X`, but `T` does not. So we can assign `&T` to `n`, but not `T{}`.

## Exploring method sets path 2

The idea of method set is **`integral to how interfaces are implemented`** and used in Go.

- **The method set of a type T consists of all methods with receiver type T**
  - These methods can be called using variables of type T.
- The Method set of a `type *T` consists of all methods with `receiver *T or T`.
  - These methods can be called using variables of type `*T`.
  - It can call methods of the corresponding non-pointer type as well.
