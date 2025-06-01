# Document

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
