package roll

import (
	"testing"
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
		number, dtype, err := ParseDice(test.input)
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

func TestRollSingleDiceInput(t *testing.T) {
	tests := []struct {
		input string
		min   int
		max   int
	}{
		{"2d6", 2, 12},
		{"1d20", 1, 20},
		{"3d8", 3, 24},
		{"5d10", 5, 50},
		{"4d12", 4, 48},
	}

	for _, test := range tests {
		dice_rolls, total, err := RollDice(test.input)
		if err != nil {
			t.Errorf("RollDice(%s) error = %v", test.input, err)
			continue
		}

		if len(dice_rolls) == 0 {
			t.Errorf("RollDice(%s) rolls map is empty", test.input)
			continue
		}
		if total < test.min || total > test.max {
			// Check if the total is greater than 0 and less than or equal to the maximum possible value
			t.Errorf("RollDice(%s) total = %d; want %d < total < %d", test.input, total, test.min, test.max)
			continue
		}
	}
}

func TestRollMultiDiceInput(t *testing.T) {
	tests := []struct {
		input string
		min   int
		max   int
	}{
		{"2d6 2d6", 4, 24},
		{"1d20 1d20", 2, 40},
		{"3d8 3d8", 6, 48},
		{"5d10 5d10", 10, 100},
		{"4d12 4d12", 8, 96},
		{"2d6 1d20", 3, 26},
		{"1d20 3d8", 4, 32},
	}

	for _, test := range tests {
		dice_rolls, total, err := RollDice(test.input)
		if err != nil {
			t.Errorf("RollDice(%s) error = %v", test.input, err)
			continue
		}

		if len(dice_rolls) == 0 {
			t.Errorf("RollDice(%s) rolls map is empty", test.input)
			continue
		}
		if total < test.min || total > test.max {
			// Check if the total is greater than 0 and less than or equal to the maximum possible value
			t.Errorf("RollDice(%s) total = %d; want %d < total < %d", test.input, total, test.min, test.max)
			continue
		}
	}
}
