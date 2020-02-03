package gedcom

import "github.com/elliotchance/gedcom/tag"

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
func Dates(nodes ...Node) (result DateNodes) {
	for _, node := range nodes {
		if IsNil(node) {
			continue
		}

		for _, n := range NodesWithTag(node, tag.TagDate) {
			result = append(result, n.(*DateNode))
		}
	}

	return
}
