## Example usage:

```go
package main

import "github.com/loginrevoked-eng/go-rust-result-enum/Result"

func main() {
    result := Result.Ok(42)
    value := result.Unwrap()
    println(value) // 42
}
```



all the methods available:
- `Ok[T any](value T) Result[T]`
- `Err[T any](err error) Result[T]`
- `IsOk() bool`
- `IsErr() bool`
- `Unwrap() T`
- `Expect(msg string) T`
- `UnwrapOr(defaultValue T) T`
- `UnwrapOrElse(f func(error) T) T`



I know this is not the best implementation of a result enum, but it's jsut for me to feel comfortable using 
Go instead of Rust.