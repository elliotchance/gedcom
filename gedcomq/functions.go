package main

// Functions is a map of available functions.
//
// See "Functions" in the package documentation for usage and examples.
var Functions = map[string]Expression{
	"?":      &QuestionMarkExpr{},
	"Length": &LengthExpr{},
}
