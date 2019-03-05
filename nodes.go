package gedcom

import (
	"reflect"
	"sync"
)

type Nodes []Node

// nodeCache is used by NodesWithTag. Even though the lookup of child tags are
// fairly inexpensive it happens a lot and its common for the same paths to be
// looked up many time. Especially when doing larger task like comparing GEDCOM
// files.
var nodeCache = &sync.Map{} // map[Node]map[Tag]Nodes{}

func NewNodes(ns interface{}) (nodes Nodes) {
	v := reflect.ValueOf(ns)
	for i := 0; i < v.Len(); i++ {
		nodes = append(nodes, v.Index(i).Interface().(Node))
	}

	return
}

// NodesWithTag returns the zero or more nodes that have a specific GEDCOM tag.
// If the provided node is nil then an empty slice will always be returned.
//
// If the node is nil the result will also be nil.
func NodesWithTag(node Node, tag Tag) (result Nodes) {
	if v1, ok1 := nodeCache.Load(node); ok1 {
		if v2, ok2 := v1.(*sync.Map).Load(tag); ok2 {
			return v2.(Nodes)
		}
	}

	defer func() {
		if v1, ok := nodeCache.Load(node); ok {
			v1.(*sync.Map).Store(tag, result)
		} else {
			nodeCache.Store(node, &sync.Map{})
		}
	}()

	if IsNil(node) {
		return nil
	}

	nodes := Nodes{}
	n := node.Nodes()
	for _, node := range n {
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
func NodesWithTagPath(node Node, tagPath ...Tag) Nodes {
	if IsNil(node) {
		return nil
	}

	if len(tagPath) == 0 {
		return Nodes{}
	}

	return nodesWithTagPath(node, tagPath...)
}

func nodesWithTagPath(node Node, tagPath ...Tag) Nodes {
	if len(tagPath) == 0 {
		return Nodes{node}
	}

	matches := Nodes{}

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

// CastTo creates a slice of a more specific node type.
//
// All Nodes must be the same type and the same as the provided t.
func (nodes Nodes) CastTo(t interface{}) interface{} {
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
	return NodesWithTag(node, tag).CastTo(t)
}

func DeleteNodesWithTag(node Node, tag Tag) {
	for _, n := range node.Nodes() {
		if n.Tag().Is(tag) {
			node.DeleteNode(n)
		}
	}
}

// FlattenAll works as Flatten with multiple inputs that are returned as a
// single slice.
//
// If any of the nodes are nil they will be ignored.
func (nodes Nodes) FlattenAll(result Nodes) {
	for _, node := range nodes {
		if IsNil(node) {
			continue
		}

		result = append(result, Flatten(node)...)
	}

	return
}

func (nodes Nodes) deleteNode(n Node) (Nodes, bool) {
	for i, node2 := range nodes {
		if node2 == n {
			return append(nodes[:i], nodes[i+2:]...), true
		}
	}

	return nodes, false
}
