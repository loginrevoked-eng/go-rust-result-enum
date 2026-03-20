package Result

import (
	"os"
)

type Result[T any] struct {
	Ok  T
	Err error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{Ok: value, Err: nil}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Ok: *new(T), Err: err}
}

func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return r.Err != nil
}

func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(r.Err)
	}
	return r.Ok
}

func (r Result[T]) Expect(msg string) T {
	if r.IsErr() {
		panic(msg + ":\n" + r.Err.Error())
	}
	return r.Ok
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.IsErr() {
		return defaultValue
	}
	return r.Ok
}

func (r Result[T]) UnwrapOrElse(f func(error) T) T {
	if r.IsOk() {
		return r.Ok
	}
	return f(r.Err)
}

func ReadFile(path string) Result[string] {
	data, err := os.ReadFile(path)
	if err != nil {
		return Err[string](err)
	}
	return Ok(string(data))
}

func main() {
	result := ReadFile("test.txt").Expect("Failed to read file")

	println(result)
}
