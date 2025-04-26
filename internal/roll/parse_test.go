package roll_test

import (
	"testing"

	"github.com/ashahide/d1c3/internal/roll"
)

func TestParseDice(t *testing.T) {
	tests := []struct {
		input  string
		number int
		dtype  int
		err    bool
	}{
		{"2d6", 2, 6, false},
		{"1d20", 1, 20, false},
		{"3d8", 3, 8, false},
		{"5d10", 5, 10, false},
		{"4d12", 4, 12, false},
		{"0d6", 0, 6, true},
		{"2d7", 2, 7, true},
	}

	for _, test := range tests {
		number, dtype, err := roll.ParseDice(test.input)
		if (err != nil) != test.err {
			t.Errorf("ParseDice(%s) error = %v; want error = %v (error: %v)", test.input, err != nil, test.err, err)
		}
		if number != test.number {
			t.Errorf("ParseDice(%s) number = %d; want number = %d", test.input, number, test.number)
		}
		if dtype != test.dtype {
			t.Errorf("ParseDice(%s) type = %d; want type = %d", test.input, dtype, test.dtype)
		}
	}
}
