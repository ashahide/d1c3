package roll

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidDieType(d int) bool {
	valid := []int{4, 6, 8, 10, 12, 20}
	for _, v := range valid {
		if d == v {
			return true
		}
	}
	return false
}

func ParseDice(dice string) (dice_number int, dice_type int, error error) {

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
