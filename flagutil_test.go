package flagutil

import (
	"flag"
	"reflect"
	"testing"
)

func TestMapString(t *testing.T) {
	m := Map{
		"hello": "world",
	}

	exp := `map[hello:world]`
	if m.String() != exp {
		t.Errorf("expected\n%v\ngot\n%v", exp, m.String())
	}
}

func TestMapSet(t *testing.T) {
	tests := []struct {
		name  string
		m     Map
		value string
		err   error
		exp   Map
	}{
		{
			name:  "bad value",
			m:     Map{},
			value: "blah",
			err:   ErrInvalidMapValue,
			exp:   Map{},
		},
		{
			name:  "spaces",
			m:     Map{},
			value: "\"h e l l o\"='w o r l d'",
			exp: Map{
				"h e l l o": "w o r l d",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.m.Set(test.value)
			if err != test.err {
				t.Errorf("expected error '%v', got '%v'", test.err, err)
			}
			if !reflect.DeepEqual(test.m, test.exp) {
				t.Errorf("expected map\n%v\ngot\n%v", test.exp, test.m)
			}
		})
	}
}

func TestFlagSetWithMap(t *testing.T) {
	// This is used to test to make sure the flag set parses correctly such that
	// we can do the actual parsing.

	args := []string{
		"-map",
		"hello=world",
		"-map",
		"test=1 2 3",
		"-map",
		"'eat'=\"cheese\"",
	}

	// Create the flag set and parse our args.
	m := Map{}
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.Var(m, "map", "a test map")
	err := fs.Parse(args)
	if err != nil {
		t.Errorf("parsing flags: %v", err)
	}

	// Verify the parse.
	exp := Map{
		"hello": "world",
		"test":  "1 2 3",
		"eat":   "cheese",
	}
	if !reflect.DeepEqual(m, exp) {
		t.Errorf("expected\n%v\ngot\n%v", exp, m)
	}
}
