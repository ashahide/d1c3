package roll_test

import (
	"testing"

	"github.com/ashahide/d1c3/internal/roll"
)

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
		dice_rolls, total, err := roll.RollDice(test.input)
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
		dice_rolls, total, err := roll.RollDice(test.input)
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
