package env

import (
	"fmt"
	"strconv"
)

// -- int Value
type intValue int

func newIntValue(val int, p *int) *intValue {
	*p = val
	return (*intValue)(p)
}

func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = intValue(v)
	return err
}

func (i *intValue) Get() interface{} { return int(*i) }

func (i *intValue) String() string { return fmt.Sprintf("%v", *i) }

// IntVar defines an int environment variable with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the variable.
func (e *Set) IntVar(p *int, name string, value int, usage string) {
	e.Var(newIntValue(value, p), name, usage)
}

// IntVar defines an int environment variable with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the variable.
func IntVar(p *int, name string, value int, usage string) {
	Env.Var(newIntValue(value, p), name, usage)
}

// Int defines an int environment variable with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the variable.
func (e *Set) Int(name string, value int, usage string) *int {
	p := new(int)
	e.IntVar(p, name, value, usage)
	return p
}

// Int defines an int environment variable with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the variable.
func Int(name string, value int, usage string) *int {
	return Env.Int(name, value, usage)
}
