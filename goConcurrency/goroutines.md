# Goroutines

Threads managed by Go runtime enabling concurrent function execution. Created with ``go`` keyword prefix. Lightweight and minimal memory overhead, can run thousands/millions cocurrently. Runtime handles scheduling across CPU cores. Communicate through channels, fundamental to Go's concurrency.

```
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

``` 
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