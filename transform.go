package gedcom

type TransformOptions struct {
	PrettyTags bool
}

func Transform(doc *Document, options TransformOptions) map[string]interface{} {
	individuals := map[string]interface{}{}
	other := []interface{}{}

	for _, node := range doc.Nodes {
		switch n := node.(type) {
		case *IndividualNode:
			individuals[n.Pointer()] = transformNode(node, options)

		default:
			other = append(other, transformNode(node, options))
		}
	}

	return map[string]interface{}{
		"individuals": individuals,
		"other":       other,
	}
}

func transformNodes(nodes []Node, options TransformOptions) []interface{} {
	ns := []interface{}{}

	for _, n := range nodes {
		ns = append(ns, transformNode(n, options))
	}

	return ns
}

func transformNode(node Node, options TransformOptions) map[string]interface{} {
	m := map[string]interface{}{}

	if options.PrettyTags {
		m["tag"] = node.Tag().String()
	} else {
		m["tag"] = string(node.Tag())
	}

	if node.Pointer() != "" {
		m["ptr"] = node.Pointer()
	}

	if node.Value() != "" {
		m["val"] = node.Value()
	}

	if len(node.Nodes()) > 0 {
		m["nodes"] = transformNodes(node.Nodes(), options)
	}

	return m
}
