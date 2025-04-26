package roll

type DiceOp struct {
	Op    string // "+" or "-"
	Value string // "2d6", "1d4", "3", etc.
	Rolls []int  // The actual rolls for this operation
	Total int    // The total of the rolls
}
