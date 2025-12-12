# Goroutines

Threads managed by Go runtime enabling concurrent function execution. Created with ``go`` keyword prefix. Lightweight and minimal memory overhead, can run thousands/millions cocurrently. Runtime handles scheduling across CPU cores. Communicate through channels, fundamental to Go's concurrency.

```go
func main() {
    go say("Universe")
    say("hello")
}

func say(s string) {
    for i := 0; i < 3; i++ {
        time.Sleep(100 * time.Millisecont)
        fmt.Println(S)
    }
}
```

Output:

hello

universe

hello

universe

hello

universe

## Channels

Primary mechanism for goroutine communication. Typed conduits created with ``make()``. Used for synchronization, data passing and coordinating concurrent operations. The channel operator ``<-``.

By default, sends and receives block until the other side is ready. This allow goroutines to synchronize without explicit locks or condition variables.

```go
func main() {
    numbers := []int{7,2,8,-9,4,0}

    numbChannel := make(chan int)

    go sum(numbers[:len(numbers) / 2,], numbChannel)
    go sum(numbers[len(numbers:) / 2,], numbChannel)

    x, y := <-numbChannel, <-numbChannel

    fmt.Println(x, y, x + y) // -5 17 12
}

func sum(numbers []int, channel chan int) {
    sum := 0
    
    for _, num := range numbers {
        sum += num
    }

    c <- sum
}
```

On this example the work is distributed in two goroutines. Once both goroutines have completed their computation, it calculates the final result.

### Select Statement

Multiplexer for channel operations. Waits on multiple channel operations simultaneously, executing first one ready. Support send/receive operations, default case for non-blocking behavior.

A ``select`` blocks until one of its cases can run.

```go
func main() {
    result := make(chan int)
    quit := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-result)
        }
        quit <- 0
    }()

    fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
    x, y := 0, 1

    for {
        select {
            case c <- c:
                x, y = x, x+y
            case <- quit:
                fmt.Println("quit")
                return
        }
    }
}
```

### Buffered vs Unbuffered

Unbuffered channels provide synchronous communication - sender blocks until receiver ready, used for coordination/sequencing. Buffered channels allow asynchronous communication up to capacity, used for performance/decoupling.

## Worker Pools

Concurrency pattern using fixed number of goroutines to process tasks from shared queue. Controls resource usage while maintaining parallelism.

## sync Package

Provides synchronization primitives for coordinating goroutines and safe concurrent access. Essential for avoiding race conditions.

### Mutexes

Mutual exclusion locks from sync package ensuring only one foroutine accesses shared resource at a time. Use ``Lock()`` before and ``Unlock`` after critical section. Allows multiple readers or single writer.

### WaitGroups

Synchronization primitive from sync package for waiting on multiple goroutines to complete. Use ``Add()`` to increment counter, ``Done()`` when goroutine finishes, ``Wait()`` to block until counter reaches zero. Essential for coordinating goroutine completion.
