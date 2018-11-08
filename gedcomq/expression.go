package main

// Expression is a single operation. Expressions can be chained together with a
// pipe (|) in the query.
type Expression interface {
	// Evaluate should only be run once and is likely to alter the value of
	// input. This means expressions can only be safely run once and previous
	// input values cannot be reused.
	Evaluate(input interface{}) (interface{}, error)
}
