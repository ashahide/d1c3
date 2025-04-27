// Package roll provides utilities for parsing, rolling, and calculating dice expressions.
package roll

import (
	"fmt"
	"math/rand"

	"github.com/ashahide/d1c3/internal/logtools"
)

// RollDice simulates rolling a list of DiceOp objects.
// It parses each dice expression, rolls the specified number and type of dice,
// stores each individual roll, and computes the subtotal for each DiceOp.
// Returns the updated list of DiceOp objects with Rolls and Total fields populated, or an error if parsing fails.
func RollDice(dice []DiceOp, advantage bool, disadvantage bool) ([]DiceOp, error) {
	logtools.Logger.Println("[RollDice] Starting dice rolling process...")
	logtools.Logger.Println("[RollDice] Dice to roll:", dice)
	logtools.Logger.Println("[RollDice] Number of dice operations to process:", len(dice))

	for i := range dice {
		d := &dice[i] // Access the real DiceOp by reference

		logtools.Logger.Printf("[RollDice] Processing dice expression: %s (operator: %s)", d.Value, d.Op)

		dice_number, dice_type, parseErr := ParseDice(d.Value)
		if parseErr != nil {
			logtools.Logger.Printf("[RollDice] Error parsing dice string '%s': %s", d.Value, parseErr)
			return nil, fmt.Errorf("error parsing dice: %w", parseErr)
		}

		total_roll := 0
		logtools.Logger.Printf("[RollDice] Rolling %d dice of type d%d", dice_number, dice_type)

		// Roll the dice the specified number of times
		for j := 0; j < dice_number; j++ {
			logtools.Logger.Printf("[RollDice] Rolling dice number: %d", j+1)

			roll := rand.Intn(dice_type) + 1 // Random roll between 1 and dice_type
			logtools.Logger.Printf("[RollDice] Rolled value: %d", roll)

			if advantage || disadvantage {
				logtools.Logger.Printf("[RollDice] Rolling second dice roll")
				roll2 := rand.Intn(dice_type) + 1 // Random roll between 1 and dice_type
				logtools.Logger.Printf("[RollDice] Rolled second value: %d", roll2)

				og_roll := roll

				if advantage {

					roll = max(roll, roll2)
					logtools.Logger.Printf("[RollDice] Advantage so choosing max(%d, %d) = %d", og_roll, roll2, roll)

				} else if disadvantage {
					roll = min(roll, roll2)
					logtools.Logger.Printf("[RollDice] Disadvantage so choosing min(%d, %d) = %d", og_roll, roll2, roll)
				}
			}

			total_roll += roll
			logtools.Logger.Printf("[RollDice] Current subtotal after roll %d: %d", j+1, total_roll)

			d.Rolls = append(d.Rolls, roll) // Append individual roll to Rolls slice
		}

		// Save the subtotal to the DiceOp
		d.Total = total_roll
		logtools.Logger.Printf("[RollDice] Final total for %s: %d", d.Value, d.Total)
	}

	logtools.Logger.Println("[RollDice] Completed rolling all dice.")
	return dice, nil
}

// GetTotal calculates the overall total of all DiceOp objects.
// It applies the operator ('+' or '-') associated with each DiceOp to the running total.
// Returns the final computed total as an integer.
func GetTotal(dice []DiceOp) (total int) {
	logtools.Logger.Println("[GetTotal] Starting total calculation...")
	logtools.Logger.Printf("[GetTotal] Number of dice operations to sum: %d", len(dice))

	for _, d := range dice {
		logtools.Logger.Printf("[GetTotal] Processing operation: %s %s (%d)", d.Op, d.Value, d.Total)

		if d.Op == "+" {
			total += d.Total
			logtools.Logger.Printf("[GetTotal] Added %d to total. New total: %d", d.Total, total)
		} else if d.Op == "-" {
			total -= d.Total
			logtools.Logger.Printf("[GetTotal] Subtracted %d from total. New total: %d", d.Total, total)
		} else {
			logtools.Logger.Printf("[GetTotal] Unknown operator '%s' encountered, ignoring.", d.Op)
		}
	}

	logtools.Logger.Printf("[GetTotal] Final calculated total: %d", total)
	return total
}
