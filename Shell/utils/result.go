package result

// Result represents either a success value or an error
type Result[T any] struct {
    value T
    err   error
}

// Ok creates a new successful Result with a value
func Ok[T any](value T) Result[T] {
    return Result[T]{
        value: value,
        err:   nil,
    }
}

// Err creates a new failed Result with an error
func Err[T any](err error) Result[T] {
    return Result[T]{
        err: err,
    }
}

// IsOk returns true if the Result contains a value
func (r Result[T]) IsOk() bool {
    return r.err == nil
}

// IsErr returns true if the Result contains an error
func (r Result[T]) IsErr() bool {
    return r.err != nil
}

// Unwrap returns the value if Result is Ok, panics if Result is Err
func (r Result[T]) Unwrap() T {
    if r.IsErr() {
        panic(r.err)
    }
    return r.value
}

// UnwrapOr returns the value if Result is Ok, returns default value if Result is Err
func (r Result[T]) UnwrapOr(defaultValue T) T {
    if r.IsErr() {
        return defaultValue
    }
    return r.value
}

// Error returns the error if Result is Err, returns nil if Result is Ok
func (r Result[T]) Error() error {
    return r.err
}