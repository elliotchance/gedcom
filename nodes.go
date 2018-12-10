package gedcom

import (
	"reflect"
	"sync"
)

// Noder allows an instance to have child nodes.
type Noder interface {
	// Nodes returns any child nodes.
	Nodes() []Node

	// AddNode will add a child to this node.
	//
	// There is no restriction on whether a node is not allow to have children
	// so you can expect that no error can occur.
	//
	// AddNode will always append the child at the end, even if there is is an
	// exact child that already exists. However, the order of node in a GEDCOM
	// file is almost always irrelevant.
	AddNode(node Node)

	// SetNodes replaces all of the child nodes.
	SetNodes(nodes []Node)
}

// nodeCache is used by NodesWithTag. Even though the lookup of child tags are
// fairly inexpensive it happens a lot and its common for the same paths to be
// looked up many time. Especially when doing larger task like comparing GEDCOM
// files.
var nodeCache = &sync.Map{} // map[Node]map[Tag][]Node{}

// NodesWithTag returns the zero or more nodes that have a specific GEDCOM tag.
// If the provided node is nil then an empty slice will always be returned.
//
// If the node is nil the result will also be nil.
func NodesWithTag(node Node, tag Tag) (result []Node) {
	if v1, ok1 := nodeCache.Load(node); ok1 {
		if v2, ok2 := v1.(*sync.Map).Load(tag); ok2 {
			return v2.([]Node)
		}
	}

	defer func() {
		if _, ok := nodeCache.Load(node); !ok {
			nodeCache.Store(node, &sync.Map{})
		}

		v1, _ := nodeCache.Load(node)
		v1.(*sync.Map).Store(tag, result)
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

// Nodes is the safer alternative to using Nodes() directly on the instance.
//
// If n is nil then nil will also be returned.
func Nodes(n Noder) []Node {
	if IsNil(n) {
		return nil
	}

	return n.Nodes()
}
