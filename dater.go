package gedcom

// Dater is implemented by nodes that are reasonably expected to have dates
// associated with them, such as events.
type Dater interface {
	Dates() DateNodes
}

// Dates returns the shallow DateNodes.
//
// Dates is safe to use with nil nodes.
//
// Dates will always return all dates, even dates that are invalid.
func Dates(node Node) (result DateNodes) {
	if IsNil(node) {
		return
	}

	for _, n := range NodesWithTag(node, TagDate) {
		result = append(result, n.(*DateNode))
	}

	return
}
