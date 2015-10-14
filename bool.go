package env

import (
	"fmt"
	"strconv"
)

// -- bool Value
type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val
	return (*boolValue)(p)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

func (b *boolValue) Get() interface{} { return bool(*b) }

func (b *boolValue) String() string { return fmt.Sprintf("%v", *b) }

// BoolVar defines a bool environment variable with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the variable.
func (e *Set) BoolVar(p *bool, name string, value bool, usage string) {
	e.Var(newBoolValue(value, p), name, usage)
}

// BoolVar defines a bool environment variable with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the variable.
func BoolVar(p *bool, name string, value bool, usage string) {
	Env.Var(newBoolValue(value, p), name, usage)
}

// Bool defines a bool environment variable with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the variable.
func (e *Set) Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	e.BoolVar(p, name, value, usage)
	return p
}

// Bool defines a bool environment variable with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the variable.
func Bool(name string, value bool, usage string) *bool {
	return Env.Bool(name, value, usage)
}
