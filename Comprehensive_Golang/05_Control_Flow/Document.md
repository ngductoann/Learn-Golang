# Document

## Go runtime, the stack and heap memory

### Go runtime

**The Go runtime system** is a component of the Go programming language that manages and scheduler the execution of Go programs. It includes a number of features that make it easy to write concurrent and parallel programs in Go.

The Go runtime system is responsible for the following

- **Goroutine management:**
  - Goroutine are lightweight threads that enable concurrent execution of Go program.
  - The runtime system manages the creation and destruction of goroutines, and schedules them for execution on available processor.
- **Garbage collection**:
  - The runtime system includes a garbage collector that automatically frees memory that is no longer needed by a program. This makes it easier to write complex programs without worrying about manual memory.
- **Memory management**:
  - The runtime system includes a memory allocator that manages the allocation and deallocation of memory used by a program.
- **Channel management**:
  - Channels are used in Go to communicate between goroutines. The runtime system manages the creation and destruction of channels, and ensures that messages are delivered in the correct order.
- **Stack management**:
  - Each goroutine has its own stack, which is used to store local variables and function arguments. The runtime system manages the allocation and deallocation of stack space for each goroutines

### The stack & The heap

In Go, the stack and heap are two regions of memory used for storing variables and data during program execution.

- **The Stack:**
  - Is a region of memory used for storing variables that are local to a function or a goroutines.
  - When a function is called or a goroutines is created, a new stack frame is created on the stack to store the function arguments, local variables, and other data.
  - The stack is a **LIFO** (last-in-firt-out) data structure, which means that the most recently added data is the first to be removed.
  - The stack is generally faster than the heap because it is managed automatically by the Go runtime system, and memory allocation and deallocation is relatively fast.

- **The heap:**
  - Is a region of memory used for storing variables that have a longer lifetime than those stored on the stack. When a variable is allocated on the stack.
  - When a variable is allocated on the heap, it remains there until it is explicitly deallocation by the program or util the program exits.
  - The heap is a more flexible data structure than the stack, but it is generally slower because it requires more overhead for memory allocation and deallocation.
  - In Go, the heap is managed automatically by the garbage collector, which frees up memory that is no longer being used by the program.

In general, Go program try to allocate as much memory as possible on the stack because it is faster and requires less overhead. However, when a program needs to store large amounts of data or data with a longer lifetime than the stack, it must use the heapa. The Go runtime system provides automatic memory management for both the stack and the heap, making it easy to write effective and scalable concurrent programs.

## The Go compiler

**The Go compiler** is a program that translates Go source code into machine-readable binary code that can be executed by a computer. In other words, the Go compiler takes the human-readable code that a programmer writes in Go and converts it into instructions that a computer can understand and execute.

The Go compiler is responsible for several important tasks, including:

- Parsing: The compiler reads the source code and breaks it down into a series of tokens, which are then analyzed for syntax errors.
- Type checking: The compiler checks that the types of variables and expressions in the code are consistent and correct according to the rules of the Go language.
- Optimizaztion: The compiler performs various optimizaztions to the code to make it run more efficiently on the target machine.
- Code generation: The compiler generates machine-readable binary code that can be executed by the target machine.

In Go, the standard compiler is called "gc" (short for "Go Compiler"). It is a highly optimized and efficient compiler that generates fast and efficient code.

## Understanding control flow

Control flow:

- sequence
- conditional
- loop

Control flow refers to the order in which statements, instructions, and operations are executed in a computer program. It determines the sequence in which instructions are executed and the conditions that determines whether or not certain instructions are executed.
In other words, control flow governs the flow of execution in a program. It is typically controlled by conditional statements(eg., if/else statements), loops (e.g., for loops, while loops), and function calls.
These constructs allow programmers to **control the flow of their programs and make decisions based on various conditions.** For example, in an if/else statement, the program executes a certain block of code if a particular condition it met and a different block of code if the condition is not met.
Similarly, in a loop, the program repeats a certain block of code unit a particular condition is met.
Overall, control flow is essential in programming because it enables programmers to create complex logic and algorithm that can perform a wide range of tasks.

## If statements & Comparison operations

And **if statements** is a programming constructs used to control the flow of execution in a program based on a condition. The basic syntax of an if statements is:

```go
if (condition){
  // Code to execute if the condition is true
}
```

In this syntax, condition is an expression that evaluates to either true or false. If the condition is true, the code inside the curly braces will be executed. If the condition is false, the code inside the curly braces will be skipped and execution will continue with the next statements following the if statement.

**Comparison operations** compare two operations and yield an untyped boolean value.

```
== equal
!= not equal
< less
<= less or equal
> greater
>= greater or equal
```

## Switch statements

A switch statements is a control flow statement in programming that allows a program to evaluate an expression or variable and then selectively execute different blocks of code based on the value of the expression or variable. It provides a convenient way to write code that performs da program to evaluate an expression or variable and then selectively execute different blocks of code based on the value of the expression or variable.
It provides a convenient way to write code that performs different actions based on a single variable or expression without having to use multiple if/else statements.
The switch statement typically consists of a single expression or variable followed by a series of case statements that specify the different values that the expression or variable can take and the corresponding blocks of code to execute in each case.
if the expression or variable matches one of the specified case values, the corresponding block of code is executed, and if none of the cases match, an optional default case can be used execute a fallback block of code.

# Using select statement from concurrency communicate

- **Concurrency** and **parallel** are related but distinct concepts in programming, including in the Go programming language.
- **Concurrency** refers to code that is written in a concurrent design pattern. This means that the code has the potential to execute multiple tasks simultaneously, where each task may make progress independently of the others.
  - In Go, concurrency is achieved using goroutines, lightweight threads of execution that are managed by the Go runtime.
  - A Go program can create many goroutines that run concurrently, each performing a different task.
  - The communication and synchronization of these goroutines is typically done using **channels** , which provide a way of goroutines to exchange data and coordinate their execution.
- **Parallelism**, on the other hand, refers to the ability of a program to execute multiple tasks simultaneously by utilizing multiple CPUs or cores.
  - Parallelism can often speed up the execution of a program by a program by allowing multiple parts of the program to run in parallel on different processors. In Go, parallelism can be achieved by running multiple goroutines on different processors using the go keyword.
    - Serial/ Sequence execution
      - The oppposite of parallel computing
      - The oppposite of code running in parallel is code running serially or sequentially. In sequential execution, each instruction or task is executed on after the other in a predefined order, so that each instruction must wait for the previous one to finish before it can start. This differs from parallel execution, where multiple instructions or tasks can be executed simultaneously. Sequential execution is typically used when the instructions or tasks are dependent on each other, or when resources required to execute the code are limited. Parallel execution, on the other hand, is used to speed up the execution of code by running multiple instructions or tasks at the same time, often on multiple processors or cores
- Go makes it easy to write concurrent code using goroutines and channels
