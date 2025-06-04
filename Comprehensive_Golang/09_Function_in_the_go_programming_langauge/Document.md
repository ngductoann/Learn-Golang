# Document

<!--toc:start-->
- [Document](#document)
  - [Introduction to functions](#introduction-to-functions)
  - [Syntax of functions in Go](#syntax-of-functions-in-go)
  - [Variadic parameter](#variadic-parameter)
  - [Unfurling a slice](#unfurling-a-slice)
  - [Defer](#defer)
  - [Methods](#methods)
  - [Interface & polymorphism](#interface-polymorphism)
  - [Exploring the stringer interface](#exploring-the-stringer-interface)
  - [Expanding on the stringer interface - wrapper func for logging](#expanding-on-the-stringer-interface-wrapper-func-for-logging)
  - [Writer interface & writing to a file](#writer-interface-writing-to-a-file)
    - [The relationship between a **STRING** and a **[]BYTE**](#the-relationship-between-a-string-and-a-byte)
  - [Writer interface & writing to a byte buffer](#writer-interface-writing-to-a-byte-buffer)
  - [Anonymous func](#anonymous-func)
  - [Func expression](#func-expression)
  - [Returning a func](#returning-a-func)
  - [Callback](#callback)
  - [Closure](#closure)
  - [Recursion](#recursion)
  - [Wrapper function](#wrapper-function)
<!--toc:end-->

## Introduction to functions

- Introduction to modular code
  - spaghetti code
  - structured / procedural programming
    - purpose
      - "abstract" code
      - code reusability
      - more understandable
    - functions
    - packages

In Go programming language, a function is a group of statements that together perform a specific task. It is an important concept as it allows us to create reusable and modular code.

Every Go program has at least one function, which is the `main()` function. The execution of a Go program starts and ends with function

**Here's the basic syntax of a function in Go**

```go
func functionname(parametername type) returntype {
  // function body
}
```

**Breakdown of the syntax:**

1. **`func`**: This is the keyword to start the declaration of a function.
2. **`functionname`**: This is the name of the function. The function name and the parameter list together constitute the function signature.
3. **`parametername type`**: These are the parameters accepted by the function. You can provide any number of parameters separated by a comma. Each parameter has a name and a type.
4. **`returntype`**: This is the data type of the value the function returns. If a function doesn't return any value, the return type can be omiited.
5. **`{function body}`**: The function body is enclosed between `{}`. The function body can contain statements and also a return statement to return a value from the function.

**For example:**

```go
func add(x int, y int) int {
  return x + y
}

func main(){
  result := add(10, 20)
  fmt.Println(result)
}
```

In this example, we've defined a function named `add` that takes two integer parameters and returns an integer. It adds the two numbers and returns the result. In the `main()` function, we call the `add` function with two numbers, and the print the result.

**Go supports multiple return values from a function. Here's how**

```go
func rectangle(length, width int) (int, int){
  var area = length * with
  var perimeter = (length * width) * 2
  return area, perimeter
}

func main(){
  area, perimeter := rectangle(10, 20)
  fmt.Println("Area:", area)
  fmt.Println("Perimeter:", perimeter)
}
```

In this example, the `rectangle` function calculates the area and perimeter of a rectangle, and return them. In the `main()` function, we call the `rectangle` function and print the returned area and perimeter.

There's a lot more to functions in Go, like variadic functions, function literals (also known as anonymous functions), higher-order functions, and more. However, this should give you a good starting point for understanding functions in Go.

## Syntax of functions in Go

- **func (receiver) identifier(parameters) (returns) { code }**
  - parameters & arguments
    - we defined our func with **parameters** (if any)
    - we call our func and pass in **arguments** (in any)
- Everything in Go is **PASS BY VALUE**
- sample func (code example main.go)
  - no params, no returns
  - 1 param, no returns
  - 1 param, 1 return
  - 2 params, 2 returns

## Variadic parameter

You can create **a func which takes an unlimited number of arguments.** When you do this, this is known as a "variadic parameter". When use the lexical element operator "...T" to signify a variadic parameter (There "T" represents some type).

- create a func that takes an unlimited number of VALUE of TYPE int
A variadic parameter in Go programming language is a parameter that can take an arbitrary number of arguments of a particular type. It's a way of allowing a function to accept as many arguments as needed.

You declare a variadic function with an ellipsis `...` before the type of the last parameter. Here's an example:

```go
func sum(nums ...int) int {
  total := 0

  for _, num := range nums {
    total += num
  }
  return total
}
```

In the `sum` function, `nums` is a variadic parameter of type `int`. The function can be called with any number of `int` arguments:

```go
sum(1, 2)
sum(1, 2, 3)
sum(1, 2, 3, 4)
```

Inside the function, `nums` is an `[]int` (slice of `int`). The `range` keyword is used to iterate over this slice.

## Unfurling a slice

When you have a slice of some type, you can **pass in the individual values in a slice by using the "..." operator**
**When you pass a slice of type T into a function that has a variadic parameter of type T in Go, you're essentially passing a variable number of arguments of the same type to the function. In Go, this feature is called "variadic functions"**

**However, it's important to note that you can't directly pass a slice to a variadic function. You need to use the ellipsis (...) operator to unpack the slice, as shown in this example:**

```go
func Sum(nums ...int) int {
  sum := 0
  for _, sum := range nums {
    sum += num
  }
  return sum
}

func main(){
  nums := []int{1,2,3,4}
  fmt.Println(Sum(nums...)) // here nums slice is unpacked using ...
}
```

In the example above, `nums...` is the use of the ellipsis operator to unpack the `nums` slice into a number of arguments to match the `nums ...int` parameter the `Sum` function. This is the mechanism used to pass a slice to a variadic function in Go.

## Defer

- A "defer" statement invokes a function whose **execution is deferred to the moment the surrounding function returns,** either because
  - The surrounding function executed a **return** statement
  - reached the **end** of its function body
  - or because the corresponding goroutine is **panicking.**

In Go programming language, the `defer` keyword is used to ensure that a function call is executed late in a program's execution, usually for purposes of cleanup. `defer` is often used when `ensure` and `finally` would be used in other languages.

When a function call is deferred, it's placed onto a stack. The function calls on the stack are executed when the surrounding function returns, not when the `defer` statement is called. They're executed in Last In, First Out (LIFO) order, so if you have multiple `defer` statements, the most recently deferred function will be the first one to execute when the function returns.

Here's a simple example:

```go
package main

import "fmt"

func main(){
  defer fmt.Println("world")
  fmt.Println("hello")
}
```

In this example, the output will be:

```text
hello
world
```

Despite the "world" line being deferred before "hello" is printed, it only actually gets printed after the `main` function has finished executing all its other statements, including printing "hello".

This makes `defer` particularly useful for handling things like closing files or releasing resources. For instance, you can open a file, defer the `Close()` operator, and then proceed with reading from or writing to the file, confident in the knowledge that it will be closed properly when the function is done, no matter happens:

```go
package main

import (
  "fmt"
  "os"
)

func main() {
  f, err := os.open("myfile.txt")
  if err != nil {
    fmt.Println(err)
    return
  }

  defer f.close()

  // continue to work with file
}
```

## Methods

**A method is nothing more than a FUNC attached to a TYPE.** When you attach a func to a type it is a method of that type. You Attach a func to a type with a RECEIVER.
Methods in the Go programming language are essentially functions, but they're associated with a specific type. The type they are associated with is called the receiver type, and the method can bee called using any variable of the type.

Here's the basic syntax for defining a method in Go:

```go
func (receiver ReceiverType) MethodName(parameters) ReturnType {
  // method body
}
```

In this syntax:

- **`func`** is the keyword that starts a function declaration.
- **`receiver`** is a variable that you will use to access the data structure in the method. It is effectively an argument of the method.
- **`ReceiverType`** is the type that this method is accosicated with. It can be any type except for interface types or pointer types.
- **`MethodName`** is the name of the method. It follows the same conventions as naming functions in Go.
- **`parameters`**  and **`ReturnType`** work just like in a function declaration. The parameters are input, and the return type is what the function will output.

Here is an example of a method in Go. Assume we have a struct type `Circle`, with a method `Area`:

```go
type Circle struct {
  Radius float64
}

func (c Circle) Area() float64 {
  return math.Pi * c.Radius * c.Radius
}
```

In this example, `Area` is a method with receiver type `Circle`. You can call it like this:

```go
c := Circle{Radius:5}
fmt.Println(c.Area())
```

In this case, `c` is an instance of type `Circle`, and `Area` is a method that can be called on instance of `Circle`.

Note: If a method needs to modify the receiver or the receiver is large, it's common to use a pointer to the receiver type instead:

```go
func (c *Circle) Scale(factor float64){
  c.Radius *= factor
}
```

In this case, `Scale` is a method with receiver type `*Circle` (pointer to Circle). It modifies the `Radius` of the circle it is called on. The method can be called as follows:

```go
c := &Circle{Radius: 5}
c.Scale(2)
fmt.Println(c.Radius) // Output: 10
```

## Interface & polymorphism

An **interface** in Go defines a set of method signature. **Polymorphism** is the ability of a VALUE of a certain TYPE to also be of another TYPE.

**In Go, values can be of more than one type.**

- **Polymorphism** refers to the ability of an object to be of an additional TYPE and take on different behaviors. In the context of programming, it allows VALUES of different TYPES to be treated as instances of a common TYPE.
- Imagine you have TYPES such as "Dog", "Cat" and "Bird"
- You can make an interface TYPE called "Animal" that requires a method called "MakeSound()" to implement this TYPE.
- Now, you can have TYPES "Dog", "Cat", and "Bird" define their own implementation of the method "MakeSound()" and any VALUES of TYPE "Dog", "Cat", and "Bird" will also be of TYPE "Animal".
- When you call the "MakeSound()" method for any VALUE of TYPE "Animal", the specific implementation for the underlying TYPE ("Dog", "Cat", or "Bird") will be invoked. This behavior is known as polymorphism.
- It simpler terms, polymorphism allows you to treat different VALUES of different TYPES as the same interface TYPE when they share a common interface. This enables you to write code that can operate on a varietyu of TYPEs without needing to know their exact underlying TYPE, as long as they adhere to the shared behavior defined by the interface TYPE.

An interface allows a value to be of more than one type.

We create an interface using this syntax: **"keyword identifier type"** so for an interface it would be: "type human interface" We then define which method(s) a type must have to implement that interface.

if a TYPE has the required method, which could be none (**the empty interface** denoted by interface{} or any), then that TYPE **implicitly implements** the interface and is **also** of that interface type.

In Go, values can be of more than one type.

1. Interface:
**An interface in Go defines a set of method signatures.** It describes the behavior or contact that a type should adhere to. Any type that implements the methods specified in an interface is said to satisfy interface. In Go, interface are **implicitly implemented;** there's no need to explicitly declare that a type implements an interfaces.

Here's an example of an interface declaration in Go:

```go
type Shape interface {
  Area() float64
  Perimeter() float64
}
```

The `Shape` interface has two methods: `Area()` and `Perimeter()`, both of which return `float64`. Any type that defines these two methods will implicitly satisfy the `Shape` interface. For instance, if we have a `Circle` type with `Area()` and `Perimeter()` methods, it satisfies the `Shape` interface.

```go
type Circle struct {
  radius float64
}

func (c Circle) Area() float64{
  return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
  return 2 * math.Pi * c.radius
}
```

Now, if we create a variable of the `Shape` type and assign a `Circle` instance to it, it is valid because `Circle` implements the `Shape` interface.

```go
var shape Shape
shape = Circle{radius: 5.0}
```

**Interface enable polymorphism in Go, allowing us to work with different types that satisfy a common interface without knowing the exact underlying type.**

2. Polymorphism:

**Polymorphism is the ability of an object to take on different forms or types.** In Go, polymorphism is achieved through interfaces. Since a type can satisfy multiple interfaces, we can use interfaces to define different behaviors for a single type.

Consider the following example:

```go
type Sayer interface {
  Say()
}

type Dog struct{ }

func (d Dog) Say(){
  fmt.Println("Woof!")
}

type Cat struct{ }

func (c Cat) Say(){
  fmt.Println("Meow!")
}

func LetThemSpeak(s Sayer){
  s.Say()
}
```

In this example, we have two types, `Dog` and `Cat`, both implementing the `Say()` method. The `LetThemSpeak()` function takes a parameter of type `Sayer`, which is an interface. When we pass a `Dog` or `Cat` instance to this function, it calls the `Say()` method specific to that instance. This behavior demonstrates polymorphism, as the function can work with different types that implement the `Sayer` interface.

```go
dog := Dog{}
cat := Cat{}

LetThemSpeak(dog) // Output: Woof!
LetThemSpeak(cat) // Output: Meow!
```

Go interfaces define sets of method signature, and types implicitly satisfy these interfaces. Polymorphism in Go is achieved through interface, allowing different types to be treated uniformly based on their shared interface implementation. This flexibility promotes code reusability and abstraction.

## Exploring the stringer interface

In the Go programming language, the `Stringer` interface is a predefined interface that **allows types to define their own string representation.** It is part the `fmt` package and consists of a single method with the signature `String() string`.

Here's the declaration of the `Stringer` interface in Go:

```go
type Stringer interface {
  String() string
}
```

To implement the `Stringer` interface, a type simply needs to define the `String()` method with the following signature:

```go
func (t T) String() string {
  // Generate and return the string representation of t
}
```

Here, 'T' represents the type implementing the `Stringer` interface. Inside the `String()` method, you should generate and return the desired string representation of the object of type `T`.

Once a type implements the `Stringer` interface, it can be used with the `fmt` package to format and print the object using the `Print` and `Printf` functions or in string interpolation. **When a value of a `Stringer` type is encountered, the `fmt` package automatically calls the `String()` method to obtain the string representation.**

Here's a simple example to illustrate the usage of the `Stringer` interface:

```go
package main

import "fmt"

type Person struct {
  Name string
  Age int
  Gender string
}

// like toString() in Java
func (p Person) String() string {
  return fmt.Sprintf("Name: %s, Age: %d, Gender: %s", p.Name, p.Age, p.Gender)
}

func main() {
  p := Person{Name: "John Doe", Age: 30, Gender: "Male"}
  fmt.Println(p) // Automatically calls the String() method of Person
}
```

In this example, the `Person` struct implements the `Stringer` interface by defining the `String()` method. The `String()` method returns a formatted string representation of the `Person` object. When `fmt.Println(p)` is called, it automatically invokes the `String()` method to obtain the string representation, resulting in the output: "Name: John Doe, Age: 30, Gender: Male".

## Expanding on the stringer interface - wrapper func for logging

- Implement a weapper func that takes type fmt.Stringer
  - log.Println

In the Go programming language, a **wrapper function**, also knows as a **wrapper**, is **a function that provides an additional layer of abstraction or functionality around an existing function** or method. It acts as an **intermediary** between the caller and the wrapped function, allowing you to modify inputs, outputs, or behavior **without directly modifying the original function.**

Wrapper functions are commonly used for various purpose, such as:

1. **Logging:** A wrapper function can add logging statements before and after invoking the wrapped function. This helps in capturing information about the function calls, input parameters, return values, and any errors that may occur.
2. **Timing and profiling:** Wrappers can be used to measure the execution time of functions, enabling performance analysis and profiling. By recording the start and end times, you can calculate the elapsed time and gater statistics.
3. **Authentication and authorization:** Wrappers can handle authentication and authorization checks before executing the wrapped function. They can validate user credentials, verify permissions, and ensure that the caller has the necessary rights to access the wrapped functionality.
4. Error handling: Wrappers can intercept errors returned by the wrapped function and transform them into a different error type or add more contextual information. They can also recover from panics and gracefully handle exceptional situations.

Here's a simple example to illustrate the concept of a wrapper function in Go:

```go
package main

import (
  "fmt"
  "time"
)

// Wrapper function for adding timing information
func TimedFunction(fn func()) {
  start := time.Now()
  fn()
  elapsed := time.Since(start)
  fmt.Println("Elapsed time: ", elapsed)
}

// Function to be wrapped
func MyFunction() {
  time.Sleep(2 * time.Second) // Simulate some work
  fmt.Println("MyFunction completed")
}

func main() {
  // Call the wrapped function
  TimedFunction(MyFunction)
}
```

In the example above, the `TimedFunction` acts as a wrapper function that measures the elapsed time taken by `MyFunction` to execute. It captures the start time, calls `MyFunction`, calculates the elapsed time, and then prints it. By using the wrapper, you can easily add timing functionality to multiple functions without modifying their implementation.

## Writer interface & writing to a file

The **writer interface** is a fundamental concept used for **writing data to various output destinations**, including files. The writer interface is defined by the `io.Writer` interface type. It specifies a single method called `Write` with the following signature:

```go
func (w Writer) Writer (p []byte) (n int, err error)
```

**The `Write` method takes a byte slice (`[]byte`) as input** and returns the number of bytes written and an error (if any). It writes the contents of the byte slice to the underlying output destination.

To write to a file using the writer interface, you need to create a file or open an existing file using the `os.Create` or `os.OpenFile` function, respectively. These functions return a file that implements the `io.Writer` interface. Here's an example:

```go
package main

import {
  "io"
  "os"
}

func main(){
  file, err := os.Create("output.txt")
  if err != nil {
    panic(err)
  }

  defer file.Close()

  data := []byte("Hello, World!")

  _, err = file.Write(data)
  if err != nil {
    panic(err)
  }
}
```

In this example, we use `os.Create` to create a new file called "output.txt" in the current directory. The `Create` function returns a file (`os.File`) the satisfies the `io.Writer` interface. We then defer the closing of the file using `file.Close()` to ensure it gets closed properly.

Next, we create a byte slice `data` the content we want to write to the file. We call `file.Write(data)` to write the byte slice to the file. The `Write` method returns the number of bytes written and any error encountered during the write operation. If there's an error, we panic and terminate the program.

by utilizing the writer interface, you can write data to files in Go using a consistent and flexible approach. Additionally, this interface is widely used in the Go standard library, enabling compatibility with various output destination beyoud files, such as network connection or other data streams.

### The relationship between a **STRING** and a **[]BYTE**

In go, a string and a []byte are two different types, but they are closely related and can often be converted between each other.

A **string** in Go represents a sequence of characters. It is an immutable type, which means you cannot modify individual characters within a string. String valuesa are always interpreted as UTF-8 encoded Unicode text.

On the other hand, a **[]byte** is a **slice of bytes**, where **each element represents a single byte.** It is a mulable type, so you can modify individual bytes within a byte slice. It can be used to represent binary data or text in various encodings.

Go provides built-in functions to convert between strings and byte slices:

1. Converting a **string to a byte slice**: You can convert a string to a byte slice using the `[]byte` type conversion. For example:

```go
str := "Hello"
bytes := []byte(str)
```

This will create a new byte slice `bytes` that contains the UTF-8 encoded representation of the string "Hello". Modifying `bytes` will not affect the original string.

2. Converting a **byte slice to a string:** You can convert a byte slice to a string using the `string()` type conversion. For example:

```go
bytes := []byte{72, 101, 108, 108, 111}
str := string(bytes)
```

This will create a new string `str` that represents the UTF-8 encoded text corresponding to the byte slice. Modifying `str` will not affect the original byte slice.

It's important to note that converting between strings and byte slices involves encoding and decoding operations, which can introduce potential errors if the encoding is not handled correctly. It's essential to be mindful of the encoding when working with these conversions to ensure accurate representation of the data.

Overall, string and byte slices provide different ways of representing and manipulating textual or binary data in Go, and Go provides convenient methods to convert between these two types when necessary.

## Writer interface & writing to a byte buffer

A byte buffer is a **region of memory** used to **temporatily store a sequence of bytes.** It provides a data structure for efficient manipulation of byte data. A byte buffer allows you to **read and write bytes to and from the buffer,** making it useful for tasks like data serialization, network communication, network communication, file I/O, and efficient string manipulation.

The concept of a byte buffer is not specific to any particular programming language but is a general concept in computer programming. Different programming languages may provide their own implementations of byte buffers, each with its own set of features and functionality.

The purpose of a byte buffer is to provide a **flexible and efficient way to work with sequences of bytes.** It typically offers methods of functions for operations such as appending data, reading data, resizing the buffer, and converting the buffer's contents to different types (e.g., strings, integers, etc.)

Byte buffer are commonly used to in scenarios where you need to manipulate raw binary data or handled I/O operations that involve byte-level operations. For example, when reading data from a network socket of a file, you can use a byte buffer to efficiently read and process chunks of bytes at a time.

A byte buffer is a data structure that provides a convenient to manipulate sequences of bytes efficiently. It serves as a temporary storage for byte data and enables operations such as reading, writing, appending, and resizing byte sequence.

In the Go programming language, `bytes.Buffer` is a type that provides a way to manipulate in-memory byte buffers. It is part of the standard library package called `bytes`.

`bytes.Buffer` is designed to efficiently handle byte manipulation tasks, such as concatenation, appending, reading, and writing to a byte buffer. It provides a convenient interface to work with sequences of bytes, similar to a file or a string.

You can create a new `bytes.Buffer` by using the `NewBuffer` function provided by the `bytes` package:

```go
import "bytes"

// Create a new bytes.Buffer
buffer := bytes.NewBuffer(nil)
```

The `nil` argument passed to `NewBuffer` initializes the buffer with an empty byte slice. Alternatively, you can initialize the buffer with an existing byte slice by passing it to `NewBuffer`, like this:

```go
data := []byte("Hello, World!")
buffer := bytes.NewBuffer(data)
```

Once you have a `bytes.Buffer`, you can perform various operations on it. Some common operations include:

- Appending data to the buffer using the `Write` or `WriteString` methods.
- Reading data from the buffer using the `Read` or `ReadString` methods.
- Getting the contents of the buffer as a byte slice using the `Bytes` method.
- Getting the contents of the buffer as a string using the `String` method.
- Resetting the buffer to its initial state using the `Reset` method.

Here's an example that demonstrates some basic operations with a `bytes.Buffer`:

```go
import (
  "fmt"
  "bytes"
)

func main() {
  buffer := bytes.NewBufferString("Hello, ")

  // Append "World!" to the buffer
  buffer.WriteString("World!")

  // Reading the contents of the buffer
  data := buffer.String()
  fmt.Println(data) // Output: Hello, World!

  // Reset the buffer
  buffer.Reset()

  // Append new data to the buffer
  buffer.WriteString("Go programming language!")

  // Read the contents of the buffer again
  data = buffer.String()
  fmt.Println(data) // Output: Go programming language!
}
```

In this example, we create a `bytes.Buffer` initialized with the string "Hello, ". We then append "World!" to the buffer, read the contents, reset the buffer, and append new data. Finally, we read the contents again to see the updated value.

`bytes.Buffer` is a flexible and efficient way to work with byte sequences in Go, often used for tasks like string manipulating, file I/O, or network communication.

## Anonymous func

**Anonymous self-executing func**

- no parameters
- 1 parameter

**An anonymous function, often also called a function literal, lambda, or closure, is a function that is defined without being given a name.**  Anonymous functions can be used in programming where a function is expected but a named function is not necessary or not desirable.

In Go, the concept of anonymous functions is supported. They can be used wherever function types are expected.
Here is an example of an anonymous function in Go:

```go
func() {
  fmt.Println("I'm an anonymous function!")
}()
```

In this example, the function is declared using the `func` keyword, it takes no arguments, and it doesn't return any value. The function's body is enclosed in `{}` and then the function is immediately invoked with `()`.

You can also assign an anonymous function to a variable and then invoke it using the variable name. Here's an example:

```go
printFunction := func() {
  fmt.Println("I'm an anonymous function assigned to a variable!")
}

printFunction() // Call the function using the variable name
```

You can also create an anonymous function that accepts arguments and returns a value. Here is example:

```go
multiply := func(a int, b, int) int {
  return a * b;
}

result := multiply(5, 10)
fmt.Println(result) // this will print 50
```

It's important to note that, since Go supports closures, an anonymous function can access and modify variables defined outside of the function itself. This is a powerful feature, but it can also lead to unexpected side if not used carefully.

## Func expression

Assigning a func to a variable

- **Assign a function to a variable**
  - execute it
- an expression is something that evaluates to a value
- "first class citizen"

**An expression** is combination of **values, variables, operations, and function calls that are evaluated to produce a single value.** It can be as simple as a literal value or a variable, or it can involve complex operations and function invocations.

Here are some examples of expressions in Go:

**1. Literal values:**

- `42`: An integer literal expression representing the value 42
- `"Hello, World!"`: A string literal expression representing the text "Hello, World!"

**2. Variables:**

- `x`: A variable expression expression representing the value stored in the variable `x`

**3. Arithmetic expressions:**

- `x + y`: An arithmetic expression that adds the values of variables `x` and `y`.
- `5 * (x - 2)`: An arithmetic expression involving multiplication, subtraction, and parentheses.

**4. Function calls:**

- `math.Sqrt(16)`: a function call expression invoking the `Sqrt` function from the `math` package with the argument `16`.
- `fmt.Sprintf("Value: %d", v)`: A function call expression invoking the `Sprintf` function from the `fmt` package, formatting the string `Value: %d` with the value of the variable `x`.

**5. Logical expressions:**

- `x > 10 && y < 20`: A logical expression combining the greater-than and less-than operations with logical AND.

**6. Type conversions:**

- `float64(x)`: A type conversion expression that converts the values of `x` to a `float64`.

**Expressions are the building blocks on Go programs, and they are used to perform computations, manipulate values, and make decisions based on conditions. They can be assigned to variables, passed as arguments to functions, or used in control flow statements like if statements and loops.**

The term **"first-class citizen"** refers to the status of certain **entities, such as values, type, and functions, that are treated equally and have the same compatibilities as other entities in the language.** It means that **these entities can be assigned to variables, passed as arguments to functions, and returned as values from functions,** just like any other data type in the language.

In Go, **functions are considered first-class citizens.** **They can be assigned to variables, passed as arguments to other functions, and returned as values from functions.** This allows Go to support high-order functions, where functions can operate on the other functions. For example, you can define a function that takes another functions as an argument and then call that function within the body of the function.

Here's an example that demonstrates the use of first-class function in Go:

```go
package main

import "fmt"

// Function that takes another function as an argument
func applyOperation(a, b int, operation func(int, int) int) int {
  return operation(a, b)
}

// Functions to be used as arguments
func add(a, b int) int {
  return a + b
}

func subtract(a, b, int){
  return a - b
}

func main(){
  result1 := applyOperation(5, 3, add) // Passing 'add' function as argument
  result2 := applyOperation(8, 4, subtract) // Passing 'subtract' function as argument

  fmt.Println(result1)
  fmt.Println(result2)
}
```

In the example, the `applyOperation` function takes two integers and a function as arguments. It applies the given function to the integers and returns the result. The `add` and `subtract` applies the given function to the integers and returns the result. The `add` and `subtract` functions are defined separately and passed as arguments to `applyOperation` when calling it.

**The ability to treat functions as first-class citizens in Go allows for more flexible and modular code, enabling the use of higher-order functions, function composition, and other functional programming techniques.**

## Returning a func

**You can return a func from a func.** Here is what that looks like.
Returning a function in the Go programming language is a form of using higher-order functions, which are functions that can accept other functions as arguments and/or return other functions as results. This is a common practice in function programming paradigms, but it's also available in multiparadigm languages like Go.

Here's an example of a function returning another function in Go:

```go
package main

import "fmt"

// This function returns another function
func incrementer() func() int {
  i := 0

  // Define and return an anonymous function
  return func() int {
    i += 1
    return i
  }
}

func main() {
  increment := incrementer()
  fmt.Println(increment()) // Output: 1
  fmt.Println(increment()) // Output: 2
  fmt.Println(increment()) // Output: 3
}
```

In this code, the `incrementer` function creates a variable `i` and then defines an anonymous function that increments `i` by one each time it's called. This anonymous function captures the `i` variable from its enclosing scope, a feature known as closure in computer science.

The `incrementer` function returns this anonymous function, and then in `main`, we call `incrementer` and assign the returned function to the `incrementer` variable. This variable is now itself a function that we can call, and each time we call it, it increments its internal `i` counter and returns the new value.

## Callback

The term **"callback"** in programming refers to **a function or piece of code that is passed as an argument to another function.** The term itself can be understood by breaking it down into "call" and "back".

The "call" part refers to the action of invoking or executing a function, when a function is called, it starts executing its code, performs certain operations, and then returns a result.
The "back" part refers to the idea that the callback function is called back or invoked by the original function after it has completed is execution or reached a specific point in its code.
Instead of the original function returning a result and ending, it "calls back" the callback function, allowing it to execute some additional code or handled specific tasks.
Callbacks are often used in event-driven programming or asynchronous operations, where the flow of execution is not linear. For example, in web development, callbacks are commonly use with JavaScript's asynchronous functions to handle responses from server requests or user interactions.
The term "callback" has become a widely accepted convention in programming to describe this mechanism of passing functions as arguments to other functions and invoking them later. It emphasizes the idea that callback function is "called back" by the original function, enhancing the understanding of the code's flow and behavior.

```go
package main

import "fmt"

// Function that takes another function as an argument
func applyOperation(a, b int, operation func (int, int) int) {
  return operation(a, b)
}

// Function to be used as arguments
func add(a, b int) int {
  return a + b
}

func subtract(a, b int) {
  return a - b
}

func main() {
  result1 := applyOperation(5, 3, add)  // Passing 'add' function as argument
  result2 := applyOperation(8, 4, subtract) // Passing 'subtract' function as argument
  fmt.Println(result1) // Output: 8
  fmt.Println(result2) // Output: 4
}
```

## Closure

- One scope enclosing other scopes
  - variables declared in the outer scope are accessible in inner scopes
- closure helps us limit the scope of variable

In programming, closure refers to the combination of a function (or a method) and the environment in which it was defined. It allows a function to access variable and bindings that are outside of its own scope, even after the outer function has finished executing.

A closure is created when a nested function references a variable from its containing (parent) function. The nested function retains a references to the environment in which in was defined, so it can access and manipulate variables from that environment, even if the parent function has already completed.

Closures ara particularly useful in situations where you want to create functions that have access to certain variables or data that persist across multiple invocations. They enable you to encapsulate data and behavior together, creating self-contained functions with their own private state.

Closures have various applications in programming, including data privacy, encapsulation, and create functions with persistent state or behavior. They can widely used in functional programming language and are a powerful tool for designing elegant and modular code.

## Recursion

- **a func that calls itself**
- factotial example

**Recursion** in programming refers to the technique of **solving a problem by breaking it down into smaller subproblems of the same type. In other word, a recursive function is a function that calls itself during its execution.**

When a recursive function is called, it solves a small instance of the same problem, and then combines the result of the smaller instance with the current instance to obtain the final result. The function continues to call itself on smaller subproblems until it reaches a base case, which is a simple case that can be solved directly without further recursion.

Recursion can be a powerful technique for solving problems that exhibit a recursive structure, such as tree traversal, graph traversal, and many mathematical calculations. However, it's important to design recursive functions carefully, ensuring that they have well-defined base cases and properly handle the termination condition, to avoid infinite recursion.

## Wrapper function

In the Go programming language, a **wrapper function** , also known as a **wrapper**, is a function that provides an provides an additional layer or abstraction or functionality around an existing function or method. It acts an **intermediary** between the caller and the wrapper function, allowing you to modify inputs, outputs, or behavior **without directly modifying the original function.** A wrapper function wraps or modifies another function's behavior.

Wrapper functions are commonly used for various purposes, such as:

1. **Logging:** A wrapper function can add logging statements before and after invoking the wrapper function. This helps in capturing information about the function calls, input parameters, return values, and any errors that may occur.
2. **Time and profiling:** Wrappers can be used to measure the execution time of functions, enabling performance analysis and profiling. By recording the start and end times, you can calculate the elapsed time and gather statistics.

3. **Authentication and authorization:** Wrappers can handle authentication and authorization checks before executing the wrapped function. They can validate user credentials, verify permissions and ensure that the caller has necessary rights to access the wrapped functionality.

4. **Error handling:** Wrappers can intercept errors returned by the wrapper function and transform them into a different error type or add more contextual information. They can also recover from panics and gracefully handle exceptional situations.

In Go programming language, a common practice is to use wrapper function for error handling. This approach is useful when you want to add more content to the error or handle different types of errors in specific way.

Here is an example of how you might create a wrapper function to handle errors when reading a file:

```go
package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func readFile(filename string) ([] byte, errors) {
  data, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, fmt.Errorf("readFile %s: %v", filename, err)
  }

  return data, nil
}

func main(){
  data, err := readFile("test.txt")
  if err != nil {
    fmt.Println("Error", err)
    os.Exit(1)
  }

  fmt.Printf("Data: %s\n", data)
}
```

In this example, the `readFile` function is a wrapper around the `ioutil.ReadFile` function. If `ioutil.ReadFile` return an error, the `readFile` function wraps this error with additional context (the name of the file that caused the error) and returns this new error.

When the `readFile` function is called from `main`, the error handling logic can use the additional context provided by the `readFile` function to provide more detailed error messages. This approach helps to create more robust and maintainable code by centralizing error handling logic within the wrapper function.

And we same this example earlier of a wrapper function in Go:

```go
package main

import (
  "fmt"
  "time"
)

// Wrapper function for adding timing information
func TimedFunction(fn func()) {
  start := time.Now()
  fn()
  elapsed := time.Since(start)
  fmt.Println("Elapsed time:", elapsed)
}

// Function to be wrapped
func MyFunction() {
  time.Sleep(2 * time.Second) // Simulate some work
  fmt.Println("MyFunction completed")
}

func main() {
  // Call the wrapped function
  TimedFunction(MyFunction)
}
```

In the example above, the `TimedFunction` acts as a wrapper function that measures the elapsed time take by `MyFunction` to execute. It captures the start time, call `MyFunction` calculates the elapsed time, and then prints it. By using the wrapper, you can easily add timing functionality to multiple functions without modifying their implementation.
