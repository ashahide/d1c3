// Package roll defines structures and functions for parsing,
// simulating, and computing dice roll expressions.
package roll

// DiceOp represents a single dice roll operation extracted from a dice expression.
// It captures the operator, the dice expression string, the individual roll results, and the subtotal.
type DiceOp struct {
	Op    string // The operator associated with this roll ("+" for addition, "-" for subtraction)
	Value string // The dice expression or number, such as "2d6", "1d4", or "3"
	Rolls []int  // A slice of integers representing individual dice roll results
	Total int    // The sum of all rolls for this operation
}
