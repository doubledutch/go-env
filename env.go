package env

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// Variable defines an environment variable
type Variable struct {
	Name     string     // name as it appears on in the environment
	Usage    string     // help message
	Value    flag.Value // value as
	DefValue string     // default value (as text)
}

// Env defines the default EnvSet
var Env = NewSet(os.Args[0])

// Set defines a set of environment variables
type Set struct {
	name      string
	parsed    bool
	variables map[string]*Variable

	output io.Writer
}

// NewSet creates a new Set
func NewSet(name string) *Set {
	return &Set{
		name: name,
	}
}

// Var sets the value, name, and usage for an environment variable in the set
func (e *Set) Var(value flag.Value, name string, usage string) {
	v := &Variable{name, usage, value, value.String()}
	_, alreadythere := e.variables[name]
	if alreadythere {
		var msg string
		if e.name == "" {
			msg = fmt.Sprintf("flag redefined: %s", name)
		} else {
			msg = fmt.Sprintf("%s flag redefined: %s", e.name, name)
		}
		fmt.Fprintln(e.out(), msg)
		panic(msg) // Happens only if flags are declared with identical names
	}
	if e.variables == nil {
		e.variables = make(map[string]*Variable)
	}
	e.variables[name] = v
}

// Parse parses the values of the environment variables in the set
func (e *Set) Parse() error {
	e.parsed = true
	for name, variable := range e.variables {
		if value := os.Getenv(name); value != "" {
			if err := variable.Value.Set(value); err != nil {
				return fmt.Errorf("invalid value %q for environment variable %s: %v", value, name, err)
			}
		} else if err := variable.Value.Set(variable.DefValue); err != nil {
			return fmt.Errorf("invalid default value %q for environment variable %s: %v", variable.DefValue, name, err)
		}

	}
	return nil
}

// Parse parses environment variables associated to the default set
func Parse() error {
	return Env.Parse()
}

// Parsed returns whether the set has already been parsed
func (e *Set) Parsed() bool {
	return e.parsed
}

// Parsed returns whether the default set has been parsed
func Parsed() bool {
	return Env.Parsed()
}

func (e *Set) out() io.Writer {
	if e.output == nil {
		return os.Stderr
	}
	return e.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (e *Set) SetOutput(output io.Writer) {
	e.output = output
}
