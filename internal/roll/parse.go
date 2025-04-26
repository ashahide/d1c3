package roll

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ashahide/d1c3/internal/logtools"
)

func isValidDieType(d int) bool {
	valid := []int{1, 4, 6, 8, 10, 12, 20}
	for _, v := range valid {
		if d == v {
			return true
		}
	}
	return false
}

// Parse a dice string like "2d6" into the number of dice and the type of dice
func ParseDice(dice string) (dice_number int, dice_type int, error error) {

	logtools.Logger.Println("[ParseDice] Parsing dice string:", dice)

	// Split dice on the 'd' character
	dice_parts := strings.Split(dice, "d")

	// Check if the dice string is valid
	if len(dice_parts) != 2 {
		return 0, 0, fmt.Errorf("invalid dice format: %s", dice)
	}
	dice_number, err := strconv.Atoi(dice_parts[0])

	if err != nil {
		return dice_number, 0, fmt.Errorf("invalid number of dice after parsing: %s", dice_parts[0])
	}

	dice_type, err = strconv.Atoi(dice_parts[1])

	if err != nil {
		return dice_number, dice_type, fmt.Errorf("invalid type of dice: %s", dice_parts[1])
	}

	// Check if the number of dice is valid
	if dice_number < 1 {
		return dice_number, dice_type, fmt.Errorf("invalid number of dice: %d", dice_number)
	}

	// Check if the type of dice is valid
	if !isValidDieType(dice_type) {
		return dice_number, dice_type, fmt.Errorf("invalid type of dice: %d", dice_number)
	}

	return dice_number, dice_type, error
}

// Parse stuff like 2d6 + 1d4 - 3
// into a map of dice rolls and a total
// Example: 2d6 + 1d4 - 3 or 2d6+1d4-3

func ParseDiceString(dice_string string) ([]DiceOp, error) {
	// Initial input
	logtools.Logger.Println("[ParseDiceString] Raw input:", dice_string)

	// Remove all whitespace
	dice_string = strings.ReplaceAll(dice_string, " ", "")
	logtools.Logger.Println("[ParseDiceString] After removing spaces:", dice_string)

	// Add a + at the beginning if missing
	if len(dice_string) > 0 && dice_string[0] != '+' && dice_string[0] != '-' {
		dice_string = "+" + dice_string
		logtools.Logger.Println("[ParseDiceString] Added leading '+':", dice_string)
	}

	var dice_ops []DiceOp
	var currentOp rune
	var currentValue strings.Builder

	for idx, r := range dice_string {
		logtools.Logger.Printf("[ParseDiceString] Index %d: Char '%c'", idx, r)

		if r == '+' || r == '-' {
			if currentValue.Len() > 0 {
				// Check if the value is a pure number (e.g., "3") — add "d1"
				currentStr := currentValue.String()
				if _, err := strconv.Atoi(currentStr); err == nil {
					currentValue.WriteString("d1")
					logtools.Logger.Printf("[ParseDiceString] Added 'd1' to pure number, new value: '%s'", currentValue.String())
				}

				logtools.Logger.Printf("[ParseDiceString] Saving op='%c' value='%s'", currentOp, currentValue.String())
				dice_ops = append(dice_ops, DiceOp{
					Op:    string(currentOp),
					Value: currentValue.String(),
				})
				currentValue.Reset()
			}
			currentOp = r
			logtools.Logger.Printf("[ParseDiceString] New operator set to '%c'", currentOp)
		} else {
			currentValue.WriteRune(r)
			logtools.Logger.Printf("[ParseDiceString] Building value: '%s'", currentValue.String())
		}
	}

	// Handle the last one after the loop
	if currentValue.Len() > 0 {
		currentStr := currentValue.String()
		if _, err := strconv.Atoi(currentStr); err == nil {
			currentValue.WriteString("d1")
			logtools.Logger.Printf("[ParseDiceString] Added 'd1' to pure number, new value: '%s'", currentValue.String())
		}

		logtools.Logger.Printf("[ParseDiceString] Final save op='%c' value='%s'", currentOp, currentValue.String())
		dice_ops = append(dice_ops, DiceOp{
			Op:    string(currentOp),
			Value: currentValue.String(),
		})
	}

	logtools.Logger.Printf("[ParseDiceString] Final parsed dice_ops: %+v", dice_ops)

	return dice_ops, nil
}
