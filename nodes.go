package gedcom

// NodesWithTag returns the zero or more nodes that have a specific GEDCOM tag.
// If the provided node is nil then an empty slice will always be returned.
func NodesWithTag(node Node, tag Tag) []Node {
	nodes := []Node{}

	if node != nil {
		for _, node := range node.Nodes() {
			if node.Tag() == tag {
				nodes = append(nodes, node)
			}
		}
	}

	return nodes
}
