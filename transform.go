package gedcom

// TransformOptions provides extra options to the Transform function. Many of
// these options are also available through CLI options on the gedcom2json
// program.
type TransformOptions struct {
	// Output tags with their descriptive name instead of their raw tag value.
	// For example, "BIRT" would be output as "Birth".
	PrettyTags bool

	// Do not include Pointer values ("ptr" attribute) in the output JSON. This
	// is useful to activate when comparing GEDCOM files that have had pointers
	// generated from different sources.
	NoPointers bool

	// Use tags (pretty or raw) as object keys rather than arrays.
	TagKeys bool

	// Convert NAME tags to a string (instead of the object parts).
	StringName bool

	// A list of tags to exclude from the output.
	ExcludeTags []Tag

	// When true only official GEDCOM tags will be included in the output.
	OnlyOfficialTags bool

	// Only output these provided tags. Leave empty to act as no filter.
	OnlyTags []Tag

	// When there are multiple names for an individual this will return the
	// first of the name nodes only.
	SingleName bool
}

func Transform(doc *Document, options TransformOptions) []interface{} {
	r := []interface{}{}

	for _, node := range doc.Nodes {
		newNode := transformNode(node, options)
		if newNode != nil {
			r = append(r, newNode)
		}
	}

	if options.TagKeys {
		r = []interface{}{reduceTagKeys(r, options)}
	}

	return r
}

func reduceTagKeys(m interface{}, options TransformOptions) interface{} {
	switch n := m.(type) {
	case []interface{}:
		r := map[string]interface{}{}

		for _, v := range n {
			tag := v.(map[string]interface{})["tag"].(string)

			if (tag == "Individual" || tag == "INDI") && !options.NoPointers {
				if r[tag] == nil {
					r[tag] = map[string]interface{}{}
				}

				ptr := v.(map[string]interface{})["ptr"].(string)
				r[tag].(map[string]interface{})[ptr] = reduceTagKeys(v, options)
				continue
			}

			if _, ok := r[tag]; ok {
				// It already exists, we may need to convert it to an array.
				if _, ok := r[tag].([]interface{}); !ok {
					// Single name.
					if tag == "Name" || tag == "NAME" && options.SingleName {
						continue
					}

					r[tag] = []interface{}{r[tag]}
				}
				r[tag] = append(r[tag].([]interface{}),
					reduceTagKeys(v, options))
			} else {
				r[tag] = reduceTagKeys(v, options)
			}
		}

		return r

	case map[string]interface{}:
		// Remove the "tag" attribute since the parent invocation has already
		// extracted it into a key.
		delete(n, "tag")

		if nodes, ok := n["nodes"]; ok {
			return reduceTagKeys(nodes, options)
		} else {
			// Remove the object wrapper.
			return n["val"]
		}

		return n
	}

	return m
}

func transformNodes(nodes []Node, options TransformOptions) []interface{} {
	ns := []interface{}{}

	for _, n := range nodes {
		newNode := transformNode(n, options)
		if newNode != nil {
			ns = append(ns, newNode)
		}
	}

	return ns
}

func transformNode(node Node, options TransformOptions) map[string]interface{} {
	// Check only.
	if len(options.OnlyTags) > 0 {
		found := false
		for _, t := range options.OnlyTags {
			if node.Tag() == Tag(t) {
				found = true
			}
		}

		if !found {
			return nil
		}
	}

	// Check excludes.
	for _, t := range options.ExcludeTags {
		if node.Tag() == Tag(t) {
			return nil
		}
	}

	if options.OnlyOfficialTags && !node.Tag().IsOfficial() {
		return nil
	}

	m := map[string]interface{}{}

	if options.PrettyTags {
		m["tag"] = node.Tag().String()
	} else {
		m["tag"] = string(node.Tag())
	}

	if node.Pointer() != "" && !options.NoPointers {
		m["ptr"] = node.Pointer()
	}

	if node.Tag() == Name && options.StringName {
		m["val"] = node.String()
	} else {
		if node.Value() != "" {
			m["val"] = node.Value()
		}

		if len(node.Nodes()) > 0 {
			m["nodes"] = transformNodes(node.Nodes(), options)
		}
	}

	return m
}
