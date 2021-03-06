package assert

import (
	"testing"
)

type Assertor struct {
	t *testing.T
}

func New(t *testing.T) Assertor {
	if t == nil {
		panic("`t` could not be nil")
	}
	return Assertor{t}
}

func (assertor Assertor) AssertThatInt(got int) Integer {
	return Integer{got: int64(got), t: assertor.t}
}

func (assertor Assertor) AssertThatString(got string) String {
	return String{got: got, t: assertor.t}
}

func (assertor Assertor) AssertThatBool(got bool) Bool {
	return Bool{got: got, t: assertor.t}
}

type function func()

func (assertor Assertor) AssertThatFunction(f function) Function {
	return Function{f: f, t: assertor.t}
}

func (assertor Assertor) AssertThatArray(arr []interface{}) Array {
	return Array{arr: arr, t: assertor.t}
}

func (assertor Assertor) AssertThatError(err error) Error {
	return Error{err: err, t: assertor.t}
}

type Integer struct {
	got int64
	t *testing.T
}

func (context Integer) IsEqualInt(expected int) {
	if context.got != int64(expected) {
		context.t.Fatalf("expected: %d, got %d", expected, context.got)
	}
}

type String struct {
	got string
	t *testing.T
}

func (context String) IsEqual(expected string) {
	if context.got != expected {
		context.t.Fatalf("expected: %s, got %s", expected, context.got)
	}
}

type Bool struct {
	got bool
	t *testing.T
}

func (context Bool) IsEqual(expected bool)  {
	if context.got != expected {
		context.t.Fatalf("expected: %t, got %t", expected, context.got)
	}
}

type Function struct {
	f function
	t *testing.T
}

func (context Function) Panics(args ... interface{})  {
	if !didPanic(context.f) {
		context.t.Fatalf("expected: panics")
	}
}

func didPanic(f function) (panicFlag bool)  {
	defer func() {
		if err := recover(); err != nil {
			panicFlag = true
		}
	}()
	f()
	return false
}

type Array struct {
	arr []interface{}
	t *testing.T
}

func (arr Array) IsEmpty() {
	if len(arr.arr) != 0 {
		arr.t.Fatalf("expected an empty slice, got %v", arr.arr)
	}
}

type Error struct {
	err error
	t *testing.T
}

func (err Error) IsNil() {
	if err.err != nil {
		err.t.Fatalf("expected nil error, got %v", err)
	}
}

func (err Error) IsNotNil() {
	if err.err == nil {
		err.t.Fatalf("expected not nil error, got nil")
	}
}

