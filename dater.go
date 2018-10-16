package gedcom

// Dater is implemented by nodes that are reasonably expected to have dates
// associated with them, such as events.
type Dater interface {
	Dates() []*DateNode
}

// Dates returns the shallow DateNodes.
//
// Dates is safe to use with nil nodes.
func Dates(node Node) (result []*DateNode) {
	if IsNil(node) {
		return
	}

	for _, n := range NodesWithTag(node, TagDate) {
		result = append(result, n.(*DateNode))
	}

	return
}
