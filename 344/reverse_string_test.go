package reverse

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	// Use map of name and struct.
	// Map's iteration order is not defined and tests will run in random
	scenarios := map[string]struct {
		input []string
		want  []string
	}{
		"hello":       {[]string{"h", "e", "l", "l", "o"}, []string{"o", "l", "l", "e", "h"}},
		"Santosh":     {[]string{"S", "a", "n", "t", "o", "s", "h"}, []string{"h", "s", "o", "t", "n", "a", "S"}},
		"even":        {[]string{"e", "v", "e", "n"}, []string{"n", "e", "v", "e"}},
		"zero length": {[]string{}, []string{}},
		"one length":  {[]string{"e"}, []string{"e"}},
	}

	for name, tc := range scenarios {
		got := ReverseString(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
