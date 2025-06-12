# Document

## Communicating Sequential Processes

Concurrency is a significant aspect of the Go programming language. It provides a way to structure a program that enables it to run more efficiently through the use of multiple processor cores and to deal with many things at once.

Concurrency in Go is achieved mainly through Goroutines and channels, which are primitives provides by the language for dealing with concurrent programming tasks.

**Goroutines**

A goroutines is a lightweight thread of execution managed by the Go runtime. It's like a thread in other language, but it's much cheaper in terms of memory and startup time. You can start a Goroutines by simply using the `go` keyword before a function call:

```go
go funcName() // funcName runs concurrently with the rest of the program
```

Goroutines have the potential to run concurrently, and they're scheduled by the Go runtime, not the operation system, which makes them very efficient. But with Goroutines alone, you can run into problems such as race conditions, where two Goroutines access shared state concurrently.

### Channels

This is where channels come in. Channels are a synchronization primitive that allow Goroutines to communicate and synchronize execution. You can send values into a channel and receive those values into another Goroutine. This allows you to coordinate tasks and send data safely between Goroutines:

```go
ch := make(chan int) // Make a channel of type int

go func() {
    ch <- 1
}

fmt.Println(<-ch) // Receive the value from the channel and print it
```

### Communicating Sequential Processes (CSP)

The concepts of Goroutines and channels are inspired by a formal language and model of computation know as Communicating Sequential Processes (CSP), developed by Sir Tony Hoare. CSP provides a way to describe patterns of interaction in concurrent systems. It's a fundamental part of the design of the Go language.

In CSP and in Go, there's no feed for locks or condition variables to manage shared state among concurrent threads or processes. Instead, data is sent between Goroutines via channels, and by default, these communication are synchronized; when a value is sent via a channel, the sending Goroutine waits until Goroutine receives the value.

This approach to concurrency is often summarized by the phrase "Do not communication by sharing memory; instead, share memory by communicating." This means it's preferred to avoid having multiple Goroutines with access to shared memory and instead use channels to pass the ownership of data between them, preventing race conditions.

By adhering to the principles of CSP, Go provides a rich and expressive framework for writing concurrent and parallel programs that are also safe and easy to reason that.

## Understanding channels

Channels Introduction

- make a channel

```go
ch := make(chan int) // Create a channel of type int
```

- putting values on a channel

```go
ch <- 42 // Send the value 42 to the channel
```

- taking values off of a channel

```go
value := <-ch // Receive a value from the channel
```

- buffered channels

```go
ch := make(chan int, 4) // Create a buffered channel with a capacity of 4
```

- channels block
  - they are like runners in a relay reace
    - they are synchronized
    - they have to pass/receive the value at the same time
      - just like runners in a relay race have to pass / receive the baton to each other at the same time
        - one runner can't pass the baton at one moment
        - and then, later, have the other runner receive the baton
        - the baton is passed/received by the runners at the same time.
      - The value is passed/received synchronously; at the same time
- channels allow us to pass values between gorountines

### Code Understand channels

#### This code doesn't work

```go
package main

import (
    "fmt"
)

func main() {
    ca := make(chan int)
    ca <- 42
    fmt.Println(<-ca)
}
```

**`IMPORTANT: CHANNELS BLOCK`**

- Channels allow
  - coordination / syncchronization / orchestration
  - buffering (buffered channels)

#### Send and Receive

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int)

    go func() {
        c <- 42 // Send the value 42 to the channel
    }()

    fmt.Println(<-c) // Receive the value from the channel and print it
}
```

#### Buffer

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int, 1)
    c <- 42
    fmt.Println(<-c)
}
```

- _`The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready.`_

**`Buffer doesn't work`**

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int, 1)
    c <- 42
    fmt.Println(<-c)

    c <- 43
    c <- 44
    fmt.Println(<-c)
}
```

**`Buffer work`**

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int, 2)
    c <- 42
    fmt.Println(<-c)

    c <- 43
    c <- 44
    fmt.Println(<-c)
    fmt.Println(<-c)
}

```

## Direction Channels

**`Channel type. Read from left to right`**

A channel provides a mechanism for `concurrently executing functions` to communicate by `sending` and `receiving` values of a specified element type. The `value` of an uninitialized channel is nil.

```text
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan") ElementType
```

The optional <- operator specifies the channel direction, send or receive. If a direction is given, the channel is directional, otherwise it is direction, otherwise it is bidirectional. A channel may be constrained only to send or only to receive by `assignment` or explicit `conversion`

```text
chan T // can be used to send and receive values of type T
chan <- float64 // can only be used to send float64
<- chan int // can only be used to receive ints
```

The <- operator associates with the leftmost chan possible:

```text
chan <- chan int // same as chan <- (chan int)
chan <- <- chan int // same as chan <- (<- chan int)
<- chan <- chan int // same as <- chan (<- chan int)
chan (<- chan int)
```

A new, initialized channel value can be made using the built-in function `make`, which takes the channel type and an optional _capacity_ as arguments:

```go
make(chan int, 100)
```

The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A nil channel is never ready for communication.

A channel may be closed with the built-in function `close`. The multi-valued assignment form of the `receive operation` reports whether a received value was sent before the channel was closed.

A single channel may be used in `send statements`, `receive operations`, and calls to the built-in functions `cap` and `len` by any number of goroutines without further synchronization. Channels act as first-in-first-out queues. For example, if one goroutine sends values on a channel and a second goroutine receives them, the values are received in the order sent.

### Code Direction Channels

- Send and Receive

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int)
    cr := make(<-chan int) // receive
    cs := make(chan<- int) // send

    fmt.Println("-----")
    fmt.Printf("c\t%T\n", c)
    fmt.Printf("cr\t%T\n", cr)
    fmt.Printf("cs\t%T\n", cs)
}
```

- **`"send and receive" means "send and receive"`**

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int)

    go func() {
        c <- 42
    }()

    fmt.Println(<-c)
}
```

- **`send means send`**

```go
package main

import (
    "fmt"
)

func main() {
    // error: "invalid operation: cs <- 42 (send to receive-only type chan<- int)"
    cs := make(chan <- int)

    go func() {
        cs <- 42 // This will cause a compile-time error because cs can only send values, not receive them
    }()

    fmt.Println(<-cs) // This will also cause a compile-time error because cs can only send values

    fmt.Printf("-------------\n")
    fmt.Printf("cs\t%T\n", cs)
}
```

- **`receive means receive`**

```go
package main

import (
    "fmt"
)

func main() {
    cr := make(<-chan int)

    go func() {
        // error: "invalid operation: cr <- 42 (send to receive-only type <-chan int)"
        cr <- 42
    }()
    fmt.Println(<-cr) // This will cause a compile-time error because cr can only receive values, not send them

    fmt.Printf("------\n")
    fmt.Printf("cr\t%T\n", cr)
}
```

`A channel may be constrained only to send or only to receive [general to specific] by conversion or assignment.`

## Using channels

- create general channels
- in funcs you can specify
  - receive channel
    - you can receive values from the channel
    - a receive channel parameter
    - in the func, you can only pull values from the channel
    - you can't close a receive channel
  - send channel
    - you can push value to the channel
    - you can't receive / pull / read from the channel
    - you can only push values to the channel

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int)         // Create a channel of type int

    go send(c)                  // Start a Goroutine to send a value to the channel
    receive(c)                  // Call the receive function to get the value from the channel

    fmt.Println("about to exit") // Print a message before exiting the program

}

// send channel
func send(c chan <- int) {
    c <- 42 // Send the value 42 to the channel
}

// receive channel
func receive(c <- chan int) {
    fmt.Println("The value received from the channel", <-c) // Receive a value from the channel and print it
}
```

**`Range stops reading from a channel when the channel is closed.`**

**`Select statements pull the value from whatever channel has a value ready to be pulled`**

building on previous code:

```go
package main

import (
    "fmt"
)

func main(){
    even := make(chan int) // Create a channel for even numbers
    odd := make(chan int)  // Create a channel for odd numbers
    quit := make(chan bool) // Create a channel to signal completion

    go send(even, odd, quit) // Start a Goroutine to send values to the channels
    receive(even, odd, quit)  // Call the receive function to get values from the channels

    fmt.Println("about to exit") // Print a message before exiting the program
}

// send channel
func send(even, odd chan<- int,quit chan<- bool ){
    for i := range 100 {
        if i%2 == 0 {
            even <- i // Send the even number to the even channel
        } else {
            odd <- i // Send the odd number to the odd channel
        }
    }

    close(even) // Close the even channel to signal no more values will be sent
    close(odd)  // Close the odd channel to signal no more values will be sent
    quit <- true // Send a signal to the quit channel to indicate completion
}

func receive(even, odd <-chan int, quit <-chan bool) {
    for {
        select {
        case v, ok := <-even: // Receive from the even channel
            if !ok { // Check if the channel is closed
                fmt.Println("Even channel closed")
                return
            }
            fmt.Println("Received from even channel:", v)
        case v, ok := <-odd: // Receive from the odd channel
            if !ok { // Check if the channel is closed
                fmt.Println("Odd channel closed")
                return
            }
            fmt.Println("Received from odd channel:", v)
        case <-quit: // Receive from the quit channel
            fmt.Println("Quit signal received")
            return
        }
    }
}

```

### Fan In (Like N to 1)

**`Taking values from many channels, and putting those values into one channel`**

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    even := make(chan int)
    odd := make(chan int)
    fanin := make(chan int)

    go send(even, odd) // Start a Goroutine to send values to the even and odd channels
    go receive(even, odd, fanin) // Start a Goroutine to receive values from the even and odd channels and send them to the fanin channel

    for v := range fanin { // Range over the fanin channel to receive values
        fmt.Println("Received from fanin channel:", v) // Print the received value
    }

    fmt.Println("about to exit") // Print a message before exiting the program
}

// send channel
func send(even, odd chain<- int) {
    for i := range 100 {
        if i%2 == 0 {
            even <- i // Send the even number to the even channel
        } else {
            odd <- i // Send the odd number to the odd channel
        }
    }
    closed(even) // Close the even channel to signal no more values will be sent
    close(odd)  // Close the odd channel to signal no more values will be sent
}

// receive channel
func receive(even, odd <- chan int, fanin chan <- int) {
    var wg sybc.WaitGroup
    wg.Add(2) // Add two to the WaitGroup for the two Goroutines

    go func() {
        for v := range even {
            fanin <- v // Send the value from the even channel to the fanin channel
        }
        wg.Done() // Signal that this Goroutine is done
    }()

    go func() {
        for v := range odd {
            fanin <- v // Send the value from the odd channel to the fanin channel
        }
        wg.Done() // Signal that this Goroutine is done
    }

    wg.Wait() // Wait for both Goroutines to finish
    close(fanin) // Close the fanin channel to signal no more values will be sent
}
```

### Fan Out (Like 1 to N)

**`Taking some work and putting the chunks of work onto may goroutines`**

```go
package main

import (
    "fmt"
    "sync"
    "math/rand"
    "time"
)

func main(){
    c1 := make(chan int) // Create a channel for the first stage of processing
    c2 := make(chan int) // Create a channel for the second stage of processing

    go populate(c1) // Start a Goroutine to populate the first channel with values
    fanOutIn(c1, c2) // Call the fanOutIn function to process values from c1 and send them to c2

    for v := range c2 { // Range over the second channel to receive processed values
        fmt.Println("Received from fanOutIn channel:", v) // Print the received value
    }

    fmt.Println("about to exit") // Print a message before exiting the program
}

func populate(c chan int) {
    for i := range 100 {
        c <- i // Send the value i to the channel
    }
    close(c) // Close the channel to signal no more values will be sent
}

func fanOutIn(c1, c2 chan int) {
    var wg sync.waitGroup

    for v := range c1 {
        wg.Add(1) // Add one to the WaitGroup for each value received
        go func(v2 int) {
            c2 <- timeConsumingWork(v2) // Send the value to the second channel after processing
            wg.Done() // Signal that this Goroutine is done
        }(v)
    }
    wg.Wait() // Wait for all Goroutines to finish
    close(c2) // Close the second channel to signal no more values will be sent
}

func timeConsumingWork(n int){
    time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
    return n + rand.Intn(1000)
}
```
