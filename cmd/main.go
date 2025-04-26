package main

import (
	"fmt"
	"os"

	"github.com/ashahide/d1c3/internal/logtools"
	"github.com/ashahide/d1c3/internal/roll"
)

type CLIArgs struct {
	DiceRolls string
}

func parseInputs() (CLIArgs, error) {
	logtools.Logger.Println("Command line arguments: ", os.Args)
	if len(os.Args) < 2 {
		logtools.Logger.Println("Missing dice roll argument - args length: ", len(os.Args))
		return CLIArgs{}, fmt.Errorf("missing dice roll argument (example: 2d6)")
	}
	logtools.Logger.Println("Dice roll argument: ", os.Args[1])
	return CLIArgs{DiceRolls: os.Args[1]}, nil
}

func main() {

	// Initialize the logger
	logtools.Initialize()

	// Parse Inputs
	logtools.Logger.Println("Parsing command line arguments")
	args, err := parseInputs()
	if err != nil {
		panic(err)
	}

	// Break the input into individual rolls
	logtools.Logger.Println("Splitting input into individual rolls")
	dice_ops, err := roll.ParseDiceString(args.DiceRolls)
	if err != nil {
		logtools.Logger.Println("Error parsing dice rolls: ", err)
		panic(err)
	}
	logtools.Logger.Println("Parsed dice rolls: ", dice_ops)

	dice_ops, err = roll.RollDice(dice_ops)
	if err != nil {
		logtools.Logger.Println("Can't parse dice rolls: ", err)
		panic(err)
	}

	fmt.Println("Dice Rolls: ", dice_ops)

	// Get the total
	total := roll.GetTotal(dice_ops)
	fmt.Println("Total: ", total)
	logtools.Logger.Println("Total: ", total)
	logtools.Logger.Println("Exiting program")

}
