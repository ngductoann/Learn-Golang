# Document

<!--toc:start-->
- [Document](#document)
  - [Struct](#struct)
    - [Introduction](#introduction)
    - [Embedded structs](#embedded-structs)
    - [Anonymous structs](#anonymous-structs)
    - [Composition](#composition)
<!--toc:end-->

## Struct

### Introduction

- A struct is an composite data type (composite data types, aka, aggregate data types, aka, complex data types).
- **Structs allow us to compose together values of different types**

### Embedded structs

- We can **take one struct and embed it in another struct** . When you do this the inner type get promoted to the outer type.

### Anonymous structs

- We can create anonymous structs also. **An anonymous struct is a struct which is not associated with a specific identifier.**

### Composition

**Composition**  data types, aka, aggregate data types, aka, **complex**  data types

In Go, **composition** refers to a way of structuring and building complex types by combining multiple simpler types. Composition is one of the fundamental principles of object-oriented programming and allow you to create more flexible and reusable code.

One of the way composition is achieved is by **embedding one struct type within another.** The fields and methods of the embedded "inner" struct become accessible to the outer struct. The inner type is said to be **promoted** to the outer type. In addition, methods of the inner type are also promoted to the outer type.

In addition, methods of the inner type are also **promoted** to the outer type. This concept is similar to inheritance in traditional object-oriented languages, but Go does not have a built-in inheritance mechainism.
Example:

```go
type Engine struct {
  // Engine fields
}

// Engine method
func (e *Engine) Start(){
  fmt.Println("Engine started")
}

type Car struct {
  Engine // Embedding the Engine struct
  // Car-specific fields
}

func main(){
  car := Car{}
  car.Start() // Calling the Start() method from the embedding Engine struct
}
```

In the example above, we have two struct types: `Engine` and `Car`. The `Car` struct embeds the `Engine` struct, which means it includes all the fields and methods of the `Engine` struct if within itself. This allows the `Car` struct to access and use the `Start()` method of the embedded `Engine` struct.

By using **composition**, **you can build complex types by combining simpler ones, promoting code reuse and modular design. It provides a way to create relationships between different types without the need for traditional inheritance.**

- It's all about type:
  - Go has OOP aspects. But it's all about **TYPE**. We create **TYPES** in Go; user-defined **TYPES.** We can then have **VALUES** of that type. We can assign **VALUES** of a user-defined **TYPE** to **VARIABLES.**

- Go is Object Oriented:
  1. Encapsulation
    a. state ("fields")
    b. behavior ("methods")
    c. exported & unexported; viewable & not viewable
  2. Reusability
    a. inheritance ("embedded types")
  3. Polymorphism
    a. interfaces
  4. Overriding
    b. "promotion"

- Traditional OOP:
  1. Classes
    a. data structure descripbing a type of object
    b. you can then create "instance"/"object" from the class/blueprint
    c. classes hold both:
      i. state / data / fields
      ii. behavior / methods
    d. public / private
  2. Inheritance
- In Go:
  1. you don't create classes, you create a **TYPE**
  2. you don't instantiate, you create a **VALUE**  of a **TYPE**
- User defined types:
  - We can declare a new type
- Named types vs anonymous types
  - **Anonymous types are indeterminate**
    - They have not been declared as a type yet
    - The compiler has flexibility with anonymous types
    - You can assign an anonymous type to a variable declared as a certain type. If the assignment can occur, the compiler will figure it out; the compiler will do an implicit conversion.
    - You cannot assign a named type to a different named type.
- Padding & architectural alignment
  - Conversion:
    - Logically organize your fields together.
    - Reusability & clarity trump performance as a design concern.
    - Go will be performant. Go for readability first. However, if you are in a situation where you need to prioritize performance: lay the fields out from largest to smallest, eg, int64, float32, bool
