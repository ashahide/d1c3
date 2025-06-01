package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ashahide/d1c3/internal/logtools"
	"github.com/ashahide/d1c3/internal/roll"
)

type CLIArgs struct {
	DiceRolls    string
	Advantage    bool
	Disadvantage bool
}

func parseInputs() (CLIArgs, error) {
	rawArgs := os.Args[1:]

	// Split: arguments before first flag
	argsBeforeFlags := []string{}
	argsFlags := []string{}
	foundFlag := false

	for _, arg := range rawArgs {
		if strings.HasPrefix(arg, "-") {
			foundFlag = true
		}

		if foundFlag {
			argsFlags = append(argsFlags, arg)
		} else {
			argsBeforeFlags = append(argsBeforeFlags, arg)
		}
	}

	logtools.Logger.Println("Args before flags:", argsBeforeFlags)
	logtools.Logger.Println("Flag arguments:", argsFlags)

	// Now parse the flags separately
	fs := flag.NewFlagSet("diceRoller", flag.ContinueOnError)
	advantage := fs.Bool("advantage", false, "Roll with advantage")
	disadvantage := fs.Bool("disadvantage", false, "Roll with disadvantage")

	if err := fs.Parse(argsFlags); err != nil {
		logtools.Logger.Println("Failed parsing flags:", err)
		return CLIArgs{}, err
	}

	logtools.Logger.Println("Advantage flag:", *advantage)
	logtools.Logger.Println("Disadvantage flag:", *disadvantage)

	if len(argsBeforeFlags) == 0 {
		return CLIArgs{}, nil
	}

	joinedStrings := strings.Join(argsBeforeFlags, " ")
	logtools.Logger.Println("Joined dice rolls:", joinedStrings)

	if *advantage && *disadvantage {
		*advantage = false
		*disadvantage = false
	}

	return CLIArgs{
		DiceRolls:    joinedStrings,
		Advantage:    *advantage,
		Disadvantage: *disadvantage,
	}, nil
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

	if (args == CLIArgs{}) {
		logtools.Logger.Println("Empty arguments - returning")
		return
	}

	// Break the input into individual rolls
	logtools.Logger.Println("Splitting input into individual rolls")
	dice_ops, err := roll.ParseDiceString(args.DiceRolls)
	if err != nil {
		logtools.Logger.Println("Error parsing dice rolls: ", err)
		panic(err)
	}
	logtools.Logger.Println("Parsed dice rolls: ", dice_ops)

	dice_ops, err = roll.RollDice(dice_ops, args.Advantage, args.Disadvantage)
	if err != nil {
		logtools.Logger.Println("Can't parse dice rolls: ", err)
		panic(err)
	}

	if args.Advantage {
		fmt.Println("Dice Rolls (with advantage): ", dice_ops)

	} else if args.Disadvantage {
		fmt.Println("Dice Rolls (with disadvantage): ", dice_ops)

	} else {
		fmt.Println("Dice Rolls: ", dice_ops)
	}

	// Get the total
	total := roll.GetTotal(dice_ops)
	fmt.Println("Total: ", total)
	logtools.Logger.Println("Total: ", total)
	logtools.Logger.Println("Exiting program")

}
