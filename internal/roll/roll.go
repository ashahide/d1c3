package roll

import (
	"fmt"
	"math/rand"
	"strings"
)

func RollDice(dice string) (dice_rolls map[string]int, total int, error error) {

	input_dice := strings.Fields(dice)

	// Initialize the dice_rolls map
	dice_rolls = make(map[string]int)

	// Initialize the total
	if len(input_dice) == 0 {
		return dice_rolls, total, fmt.Errorf("no dice rolls provided")
	}

	for _, d := range input_dice {
		dice_number, dice_type, err := ParseDice(d)
		if err != nil {
			return dice_rolls, total, fmt.Errorf("error parsing dice: %s", err)

		} else {

			// Roll dice
			total_roll := 0
			for range dice_number {
				// Roll the dice
				total_roll += rand.Intn(dice_type) + 1

				dice_rolls[d] = total_roll

			}

		}

	}

	// Calculate the total
	for _, roll := range dice_rolls {
		total += roll
	}
	// Check if the total is valid
	if total <= 0 {
		return dice_rolls, total, fmt.Errorf("invalid total: %d", total)
	}

	return dice_rolls, total, error
}
