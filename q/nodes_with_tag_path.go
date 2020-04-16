package q

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"reflect"
)

// NodesWithTagPathExpr is a function. See Evaluate.
type NodesWithTagPathExpr struct{}

// NodesWithTagPath returns all of the nodes that have an exact tag path. The
// number of nodes returned can be zero and tag must match the tag path
// completely and exactly.
//
// If the node is nil the result will also be nil.
//
// Find all Death nodes that belong to all individuals:
//
//   .Individuals | NodesWithTagPath("DEAT")
//
// From the individuals find all the Date nodes within only the Birth nodes.
//
//   .Individuals | NodesWithTagPath("BIRT", "DATE")
//
// Combine all of the birth and death dates:
//
//   Births are .Individuals | NodesWithTagPath("BIRT", "DATE") | {type: "birth", date: .String};
//   Deaths are .Individuals | NodesWithTagPath("DEAT", "DATE") | {type: "death", date: .String};
//   Combine(Births, Deaths)
//
func (e *NodesWithTagPathExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	in := reflect.ValueOf(input)

	if input == nil || in.IsNil() {
		return gedcom.Nodes(nil), nil
	}

	// Convert into a slice if needed.
	if in.Kind() != reflect.Slice {
		s := reflect.MakeSlice(reflect.SliceOf(in.Type()), 1, 1)
		s.Index(0).Set(in)
		in = reflect.ValueOf(s.Interface())
	}

	// Process args.
	var argValues []gedcom.Tag
	for _, arg := range args {
		argValue, err := arg.Evaluate(engine, nil)
		if err != nil {
			return gedcom.Nodes(nil), err
		}

		stringForTag, ok := argValue.(string)
		if !ok {
			stringForTag = fmt.Sprintf("%s", argValue)
		}

		argValues = append(argValues, gedcom.TagFromString(stringForTag))
	}

	// Process slice input.
	var results gedcom.Nodes
	for i := 0; i < in.Len(); i++ {
		node := in.Index(i).Interface().(gedcom.Node)
		results = append(results, gedcom.NodesWithTagPath(node, argValues...)...)
	}

	return results, nil
}
