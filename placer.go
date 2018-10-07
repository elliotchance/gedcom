package gedcom

type Placer interface {
	Places() []*PlaceNode
}

// Places returns the shallow PlaceNodes for each of the input nodes.
//
// Dates is safe to use with nil nodes.
func Places(nodes ...Node) (result []*PlaceNode) {
	for _, node := range nodes {
		if IsNil(node) {
			continue
		}

		for _, n := range NodesWithTag(node, TagPlace) {
			result = append(result, n.(*PlaceNode))
		}
	}

	return
}
