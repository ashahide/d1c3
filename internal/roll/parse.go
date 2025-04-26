// Package roll provides utilities to parse dice roll expressions
// and compute their results for standard tabletop gaming dice.
package roll

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ashahide/d1c3/internal/logtools"
)

// isValidDieType checks if a given die type is allowed.
// Valid die types include 1, 4, 6, 8, 10, 12, and 20.
func isValidDieType(d int) bool {
	valid := []int{1, 4, 6, 8, 10, 12, 20}
	for _, v := range valid {
		if d == v {
			return true
		}
	}
	return false
}

// ParseDice parses a dice string such as "2d6" into its number of dice and die type.
// Returns the number of dice, type of dice, and an error if the format is invalid.
func ParseDice(dice string) (dice_number int, dice_type int, error error) {
	logtools.Logger.Println("[ParseDice] Parsing dice string:", dice)

	// Split the dice string on the 'd' character
	dice_parts := strings.Split(dice, "d")

	// Validate dice format
	if len(dice_parts) != 2 {
		logtools.Logger.Println("[ParseDice] Invalid dice format, expected two parts separated by 'd'")
		return 0, 0, fmt.Errorf("invalid dice format: %s", dice)
	}

	// Parse number of dice
	dice_number, err := strconv.Atoi(dice_parts[0])
	if err != nil {
		logtools.Logger.Println("[ParseDice] Invalid number of dice:", dice_parts[0])
		return 0, 0, fmt.Errorf("invalid number of dice after parsing: %s", dice_parts[0])
	}

	// Parse die type
	dice_type, err = strconv.Atoi(dice_parts[1])
	if err != nil {
		logtools.Logger.Println("[ParseDice] Invalid type of dice:", dice_parts[1])
		return dice_number, 0, fmt.Errorf("invalid type of dice: %s", dice_parts[1])
	}

	// Validate number of dice
	if dice_number < 1 {
		logtools.Logger.Println("[ParseDice] Invalid number of dice:", dice_number)
		return dice_number, dice_type, fmt.Errorf("invalid number of dice: %d", dice_number)
	}

	// Validate die type
	if !isValidDieType(dice_type) {
		logtools.Logger.Println("[ParseDice] Invalid die type:", dice_type)
		return dice_number, dice_type, fmt.Errorf("invalid type of dice: %d", dice_type)
	}

	logtools.Logger.Printf("[ParseDice] Successfully parsed: %d dice of d%d", dice_number, dice_type)
	return dice_number, dice_type, nil
}

// ParseDiceString parses a full dice expression such as "2d6+1d4-3" into a slice of DiceOp structs.
// Each DiceOp contains an operator ('+' or '-') and a dice value (like "2d6" or "3d1").
// Single numbers are automatically treated as "Nd1" (e.g., "3" becomes "3d1").
func ParseDiceString(dice_string string) ([]DiceOp, error) {
	logtools.Logger.Println("[ParseDiceString] Raw input:", dice_string)

	// Remove all whitespace from the input
	dice_string = strings.ReplaceAll(dice_string, " ", "")
	logtools.Logger.Println("[ParseDiceString] After removing spaces:", dice_string)

	// Ensure the dice string starts with an operator
	if len(dice_string) > 0 && dice_string[0] != '+' && dice_string[0] != '-' {
		dice_string = "+" + dice_string
		logtools.Logger.Println("[ParseDiceString] Added leading '+':", dice_string)
	}

	var dice_ops []DiceOp
	var currentOp rune
	var currentValue strings.Builder

	// Iterate through the dice string character by character
	for idx, r := range dice_string {
		logtools.Logger.Printf("[ParseDiceString] Index %d: Char '%c'", idx, r)

		if r == '+' || r == '-' {
			// When encountering a new operator, save the current dice term
			if currentValue.Len() > 0 {
				currentStr := currentValue.String()
				logtools.Logger.Printf("[ParseDiceString] Completed value before operator: '%s'", currentStr)

				// If the current value is purely a number, treat it as "Nd1"
				if _, err := strconv.Atoi(currentStr); err == nil {
					currentValue.WriteString("d1")
					logtools.Logger.Printf("[ParseDiceString] Converted pure number to dice format: '%s'", currentValue.String())
				}

				logtools.Logger.Printf("[ParseDiceString] Saving DiceOp: op='%c', value='%s'", currentOp, currentValue.String())
				dice_ops = append(dice_ops, DiceOp{
					Op:    string(currentOp),
					Value: currentValue.String(),
				})
				currentValue.Reset()
			}

			currentOp = r
			logtools.Logger.Printf("[ParseDiceString] New operator set to '%c'", currentOp)
		} else {
			// Build up the current dice term
			currentValue.WriteRune(r)
			logtools.Logger.Printf("[ParseDiceString] Building current value: '%s'", currentValue.String())
		}
	}

	// After finishing the loop, save the last pending dice term
	if currentValue.Len() > 0 {
		currentStr := currentValue.String()
		logtools.Logger.Printf("[ParseDiceString] Final value after loop: '%s'", currentStr)

		if _, err := strconv.Atoi(currentStr); err == nil {
			currentValue.WriteString("d1")
			logtools.Logger.Printf("[ParseDiceString] Converted pure number to dice format: '%s'", currentValue.String())
		}

		logtools.Logger.Printf("[ParseDiceString] Saving final DiceOp: op='%c', value='%s'", currentOp, currentValue.String())
		dice_ops = append(dice_ops, DiceOp{
			Op:    string(currentOp),
			Value: currentValue.String(),
		})
	}

	logtools.Logger.Printf("[ParseDiceString] Final parsed DiceOps: %+v", dice_ops)
	return dice_ops, nil
}
