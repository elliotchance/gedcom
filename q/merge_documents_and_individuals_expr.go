package q

import (
	"errors"
	"fmt"
	"github.com/elliotchance/gedcom"
)

// MergeDocumentsAndIndividualsExpr is a function. See Evaluate.
type MergeDocumentsAndIndividualsExpr struct{}

// Evaluate merges two documents while also merging similar individuals.
func (e *MergeDocumentsAndIndividualsExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	if len(args) != 2 {
		return nil, errors.New("MergeDocumentsAndIndividuals must take two arguments")
	}

	documents := []*gedcom.Document{}

	for argNumber, arg := range args {
		value, err := arg.Evaluate(engine, nil)
		if err != nil {
			return nil, err
		}

		if doc, ok := value.(*gedcom.Document); ok {
			documents = append(documents, doc)
		} else {
			return nil, fmt.Errorf(
				"argument %d of MergeDocumentsAndIndividuals is not a Document",
				argNumber+1)
		}
	}

	mergeFn := gedcom.EqualityMergeFunction
	options := gedcom.NewIndividualNodesCompareOptions()

	return gedcom.MergeDocumentsAndIndividuals(documents[0], documents[1], mergeFn, options)
}
