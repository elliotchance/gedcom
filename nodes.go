package gedcom

import "reflect"

// nodeCache is used by NodesWithTag. Even though the lookup of child tags are
// fairly inexpensive it happens a lot and its common for the same paths to be
// looked up many time. Especially when doing larger task like comparing GEDCOM
// files.
var nodeCache = map[Node]map[Tag][]Node{}

// NodesWithTag returns the zero or more nodes that have a specific GEDCOM tag.
// If the provided node is nil then an empty slice will always be returned.
//
// If the node is nil the result will also be nil.
func NodesWithTag(node Node, tag Tag) (result []Node) {
	if v, ok := nodeCache[node][tag]; ok {
		return v
	}

	defer func() {
		if _, ok := nodeCache[node]; !ok {
			nodeCache[node] = map[Tag][]Node{}
		}

		nodeCache[node][tag] = result
	}()

	if IsNil(node) {
		return nil
	}

	nodes := []Node{}
	for _, node := range node.Nodes() {
		if node.Tag().Is(tag) {
			nodes = append(nodes, node)
		}
	}

	return nodes
}

// NodesWithTagPath return all of the nodes that have an exact tag path. The
// number of nodes returned can be zero and tag must match the tag path
// completely and exactly.
//
//   birthPlaces := NodesWithTagPath(individual, TagBirth, TagPlace)
//
// If the node is nil the result will also be nil.
func NodesWithTagPath(node Node, tagPath ...Tag) []Node {
	if IsNil(node) {
		return nil
	}

	if len(tagPath) == 0 {
		return []Node{}
	}

	return nodesWithTagPath(node, tagPath...)
}

func nodesWithTagPath(node Node, tagPath ...Tag) []Node {
	if len(tagPath) == 0 {
		return []Node{node}
	}

	matches := []Node{}

	for _, next := range NodesWithTag(node, tagPath[0]) {
		matches = append(matches, nodesWithTagPath(next, tagPath[1:]...)...)
	}

	return matches
}

// HasNestedNode checks if node contains lookingFor at any depth. If node and
// lookingFor are the same false is returned. If either node or lookingFor is
// nil then false is always returned.
//
// Nodes are matched by reference, not value so nodes that represent exactly the
// same value will not be considered equal.
func HasNestedNode(node Node, lookingFor Node) bool {
	if node == nil || lookingFor == nil {
		return false
	}

	for _, node := range node.Nodes() {
		if node == lookingFor || HasNestedNode(node, lookingFor) {
			return true
		}
	}

	return false
}

// CastNodes creates a slice of a more specific node type.
//
// All Nodes must be the same type and the same as the provided t.
func CastNodes(nodes []Node, t interface{}) interface{} {
	size := len(nodes)
	nodeType := reflect.TypeOf(t)
	sliceType := reflect.SliceOf(nodeType)
	slice := reflect.MakeSlice(sliceType, size, size)

	for i, node := range nodes {
		value := reflect.ValueOf(node)
		slice.Index(i).Set(value)
	}

	return slice.Interface()
}

func castNodesWithTag(node Node, tag Tag, t interface{}) interface{} {
	return CastNodes(NodesWithTag(node, tag), t)
}
