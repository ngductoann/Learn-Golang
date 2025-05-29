# Document

## Introduct to grouping values

- Array:

  - **a numbered sequence of VALUES of the same TYPE**
  - does not change in size
  - **used for Go internals; generally not used in your code**
  - [Array type](https://go.dev/ref/spec#Array_types)

- Slice:

  - built on top of an array
  - **holds VALUES of the same TYPE**
  - changes in size
  - has a length and a capacity
  - [Slice types](https://go.dev/ref/spec#Slice_types)

- Map:
  - **key / value storage**
  - un **unordered** group of VALUEs of one TYPE, called the element type, indexed by a set of unique keys of another type, called the key type.
  - [Map types](https://go.dev/ref/spec#Map_types)

- Struct:
  - a data structure
  - a composite / aggregate
  - allows us to **collect values of different types together**
  - [Struct types](https://go.dev/ref/spec#Struct_types)

## Array - an introduction to arrays

Array are mostly used as a building block in the Go programming language. In some instance, we might use an array, but mostly the recommendation is to **use slices instead.**

## Slice

### Composite literal

- A SLICE holds VALUES of the same TYPE. If I wanted to store all of my favorite numbers, I would use a slice of int. If I want to store all of my favorite foods, I would use a slice of string.
- We will **use a SLICE LITERAL to create a slice.** A slice literal is created by having the TYPE followed by CURLY BRACES and then putting the appropriate values in the curly brace area.

### Append to a slice

- **To append values to a slice, we use the special built in function append.** This function returns a slice of the same type.

### Slicing a slice

- We can slice a slice which means that we can **cat parts of the slice away.** We do this will the colon operator.

### Deleting from a slice

We can **delete from a slice using both append and slicing.**

### Make

- Slices are built on top of arrays. A slice is dynamic in that it will grow in size. The underlying array, however, does not grow in size.
- When we create a slice, **we can use the built in function make to specify how large our slice should be and also how large the underlying array should be.** This can enhance performance a little bit:
  - make([]T, length, capacity)
  - make([]int, 50, 100)

### Multidimensional slice

- A multi-dimensional slice is like a spreadsheet. You can have **a slice of a slice of some type.**
