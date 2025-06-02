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
	Verbose      bool
}

func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
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
	verbose := fs.Bool("verbose", false, "Verbose output")

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
		Verbose:      *verbose,
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

	if args.Verbose {

		// Determine roll context
		var label string
		switch {
		case args.Advantage:
			label = "Dice Rolls (with advantage)"
		case args.Disadvantage:
			label = "Dice Rolls (with disadvantage)"
		default:
			label = "Dice Rolls"
		}

		fmt.Println(label + ":")

		for _, op := range dice_ops {
			verb := "added     "
			if op.Op == "-" {
				verb = "subtracted"
			}

			fmt.Printf("  %s %s -> %v -> subtotal: %d\n",
				titleCase(verb),
				op.Value,
				op.Rolls,
				op.Total,
			)
		}

		// Get the total
		total := roll.GetTotal(dice_ops)
		fmt.Println("Total:", total)

		logtools.Logger.Println("Total: ", total)

	} else {
		total := roll.GetTotal(dice_ops)
		fmt.Println(total)
	}

	logtools.Logger.Println("Exiting program")

}
