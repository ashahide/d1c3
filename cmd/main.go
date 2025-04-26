package main

import (
	"fmt"
	"os"

	"github.com/ashahide/d1c3/internal/roll"
)

type CLIArgs struct {
	DiceRolls string
}

func parser() (CLIArgs, error) {
	if len(os.Args) < 2 {
		return CLIArgs{}, fmt.Errorf("missing dice roll argument (example: 2d6)")
	}
	return CLIArgs{DiceRolls: os.Args[1]}, nil
}

func main() {

	args, err := parser()
	if err != nil {
		panic(err)
	}

	fmt.Println(args.DiceRolls)

	dice_rolls, total, err := roll.RollDice(args.DiceRolls)
	if err != nil {
		panic(err)
	}

	fmt.Println("Dice Rolls: ", dice_rolls)
	fmt.Println("Total: ", total)

}
