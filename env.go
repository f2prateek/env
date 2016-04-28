package env

import (
	"os"
	"strconv"

	"github.com/f2prateek/go-pointers"
)

// Bool defines a bool flag with specified name and default value. The return
// value is the address of a bool variable that stores the value of the flag.
func Bool(name string, value *bool, usage string) *bool {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	b, err := strconv.ParseBool(env)
	if err != nil {
		return value
	}
	return pointers.Bool(b)
}

// String defines a string flag with specified name and default value. The
// return  value is the address of a string variable that stores the value of
// the flag.
func String(name string, value *string) *string {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	return pointers.String(env)
}
