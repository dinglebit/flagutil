// Package flagutil contains utility functions for the standard library flag
// package.
package flagutil

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrInvalidMapValue is returned when a value is parsed that doesn't follow
	// the key=value format.
	ErrInvalidMapValue = errors.New("invalid map value: string doesn't follow format key=value")
)

// Map implements the flag.Value interface. It allows you to pass a flag
// multiple times to a program and store it's value in a map. The flag value
// should follow the format key=value. The value can use ' and "" to make allow
// the value to contain spaces.
type Map map[string]string

// String implements the flag.Value interface.
func (m Map) String() string {
	return fmt.Sprintf("%v", map[string]string(m))
}

// Set implements the flag.Value interface.
func (m Map) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return ErrInvalidMapValue
	}
	m[strings.Trim(parts[0], `'"`)] = strings.Trim(parts[1], `'"`)
	return nil
}
