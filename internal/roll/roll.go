package roll

import (
	"fmt"
	"math/rand"

	"github.com/ashahide/d1c3/internal/logtools"
)

func RollDice(dice []DiceOp) ([]DiceOp, error) {
	logtools.Logger.Println("[RollDice] Rolling dice...")
	logtools.Logger.Println("[RollDice] Dice to roll:", dice)
	logtools.Logger.Println("[RollDice] Number of dice to roll:", len(dice))

	for i := range dice {
		d := &dice[i] // <---- grab pointer to real DiceOp

		logtools.Logger.Println("[RollDice] Rolling dice:", d.Value)

		dice_number, dice_type, parseErr := ParseDice(d.Value)
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing dice: %w", parseErr)
		}

		total_roll := 0
		logtools.Logger.Printf("[RollDice] Rolling %d d%d", dice_number, dice_type)

		for j := 0; j < dice_number; j++ {
			logtools.Logger.Printf("[RollDice] Rolling dice number: %d", j+1)

			roll := rand.Intn(dice_type) + 1
			logtools.Logger.Printf("[RollDice] Rolled: %d", roll)

			total_roll += roll
			logtools.Logger.Printf("[RollDice] Current total: %d", total_roll)

			d.Rolls = append(d.Rolls, roll) // Now modifies real d
		}

		d.Total = total_roll // Correctly saved into slice
		logtools.Logger.Printf("[RollDice] Final total for %s: %d", d.Value, d.Total)
	}

	return dice, nil
}

func GetTotal(dice []DiceOp) (total int) {
	logtools.Logger.Println("[GetTotal] Calculating total...")
	for _, d := range dice {
		logtools.Logger.Printf("[GetTotal] Adding %s: %d", d.Value, d.Total)
		if d.Op == "+" {
			total += d.Total
		} else if d.Op == "-" {
			total -= d.Total
		}
	}
	logtools.Logger.Printf("[GetTotal] Total: %d", total)
	return total
}
