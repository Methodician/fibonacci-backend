package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFibonacciLoop(t *testing.T) {
	// maybe if I test these as strings instead of arrays...
	// f := make([]int, n+1, n+2)
	resultof4 := "[0 1 1 2]"
	resultof13 := "[0 1 1 2 3 5 8 13 21 34 55 89 144]"
	// no time to figure this out...
	// resultof0 :=
	// resultOf60001 :=

	cases := []struct {
		in, want string
	}{
		{"4", resultof4},
		{"13", resultof13},
	}
	for _, c := range cases {
		digits, err := strconv.Atoi(c.in)
		if err != nil {
			t.Errorf("Recieved %q when converting input string to int", err)
		} else {
			res := FibonacciLoop(digits)
			str := fmt.Sprint(res)
			got := str
			if got != c.want {
				t.Errorf("FibonacciLoop(%q) == %q, want %q", c.in, got, c.want)
			}
		}
	}
}
