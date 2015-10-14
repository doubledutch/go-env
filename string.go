package env

import "fmt"

// -- string Value
type stringValue string

func newStringValue(val string, p *string) *stringValue {
	*p = val
	return (*stringValue)(p)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

func (s *stringValue) Get() interface{} { return string(*s) }

func (s *stringValue) String() string { return fmt.Sprintf("%s", *s) }

// StringVar defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the variable.
func (e *Set) StringVar(p *string, name string, value string, usage string) {
	e.Var(newStringValue(value, p), name, usage)
}

// StringVar defines a string environment variable with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the variable.
func StringVar(p *string, name string, value string, usage string) {
	Env.Var(newStringValue(value, p), name, usage)
}

// String defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the variable.
func (e *Set) String(name string, value string, usage string) *string {
	p := new(string)
	e.StringVar(p, name, value, usage)
	return p
}

// String defines a string environment variable with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the variable.
func String(name string, value string, usage string) *string {
	return Env.String(name, value, usage)
}
