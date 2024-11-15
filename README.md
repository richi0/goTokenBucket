# goTokenBucket [![GoDoc](https://pkg.go.dev/badge/goTokenBucket.svg)](https://pkg.go.dev/github.com/richi0/goTokenBucket)

goTokenBucket provides a simple token bucket implementation. This can be used to rate limit requests.

## What does goTokenBucket do?

goTokenBucket allows you to create a token bucket with a specified capacity and refill rate, request tokens in a blocking or non-blocking manner, and check the number of available tokens.

## How do I use goTokenBucket?

### Install

```bash
go get -u github.com/richi0/goTokenBucket
```

### Example

```go
import (
    "fmt"
    "time"
    "github.com/richi0/goTokenBucket"
)

func main() {
    tb := goTokenBucket.NewTokenBucket(10, 100, 1, 10)
    fmt.Println(tb.AvailableTokens()) // Output: 10
    tb.RequestTokenBlocking()
    fmt.Println(tb.AvailableTokens()) // Output: 9
    time.Sleep(110 * time.Millisecond)
    fmt.Println(tb.AvailableTokens()) // Output: 10
    if tb.RequestTokenNonBlocking() {
        fmt.Println("Token acquired")
    } else {
        fmt.Println("No token available")
    }
}
```

## Documentation

Find the full documentation of the package here: https://pkg.go.dev/github.com/richi0/goTokenBucket
