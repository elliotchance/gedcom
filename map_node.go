package gedcom

import "github.com/elliotchance/gedcom/tag"

// MapNode pertains to a representation of measurements usually presented in a
// graphical form.
//
// New in Gedcom 5.5.1.
type MapNode struct {
	*SimpleNode
}

// NewMapNode creates a new MAP node.
func NewMapNode(value string, children ...Node) *MapNode {
	return &MapNode{
		newSimpleNode(tag.TagMap, value, "", children...),
	}
}

// Latitude is the value specifying the latitudinal coordinate of the place
// name.
//
// The latitude coordinate is the direction North or South from the equator in
// degrees and fraction of degrees carried out to give the desired accuracy.
//
// For example: 18 degrees, 9 minutes, and 3.4 seconds North would be formatted
// as N18.150944.
//
// Minutes and seconds are converted by dividing the minutes value by 60 and the
// seconds value by 3600 and adding the results together. This sum becomes the
// fractional part of the degree’s value.
func (node *MapNode) Latitude() *LatitudeNode {
	n := First(NodesWithTag(node, tag.TagLatitude))

	if IsNil(n) {
		return nil
	}

	return n.(*LatitudeNode)
}

// Longitude is the value specifying the longitudinal coordinate of the place
// name.
//
// The longitude coordinate is degrees and fraction of degrees east or west of
// the zero or base meridian coordinate.
//
// For example: 168 degrees, 9 minutes, and 3.4 seconds East would be formatted
// as E168.150944.
func (node *MapNode) Longitude() *LongitudeNode {
	n := First(NodesWithTag(node, tag.TagLongitude))

	if IsNil(n) {
		return nil
	}

	return n.(*LongitudeNode)
}
