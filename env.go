package env

import (
	"os"
	"strconv"

	"github.com/f2prateek/go-pointers"
)

func Bool(name string, value *bool) *bool {
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

func String(name string, value *string) *string {
	env := os.Getenv(name)
	if env == "" {
		return value
	}
	return pointers.String(env)
}
