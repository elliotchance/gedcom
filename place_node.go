package gedcom

import "strings"

// PlaceNode represents a jurisdictional name to identify the place or location
// of an event.
//
// The original specification was derived from
// http://wiki-en.genealogy.net/GEDCOM/PLAC-Tag
type PlaceNode struct {
	*SimpleNode
}

// NewPlaceNode creates a new PLAC node.
//
// The original specification was derived from
// http://wiki-en.genealogy.net/GEDCOM/PLAC-Tag
func NewPlaceNode(value string, children ...Node) *PlaceNode {
	return &PlaceNode{
		newSimpleNode(TagPlace, value, "", children...),
	}
}

// JurisdictionalName determines the best name to use to identify the
// jurisdiction entities.
//
// If Format() is not empty then this is the best jurisdictional name to use
// because it should be in the form of "Name,County,State,Country". Otherwise it
// will have to fallback to using the PlaceNode value.
//
// The PlaceNode value does not have to follow any standards but it's common for
// it to have the same value as the Format would have had.
//
// If the jurisdictional name is not exactly in the form of
// "Name,County,State,Country" (including more or less components) then it
// cannot be split. In this case the Name component is the entire name and the
// County, State and Country will be empty.
func (node *PlaceNode) JurisdictionalName() string {
	name := Value(node.Format())

	if name == "" {
		name = Value(node)
	}

	return name
}

// Name is the first part of the JurisdictionalName().
//
// If the JurisdictionalName is not in the exact form of
// "Name,County,State,Country" then Name will return the entire jurisdictional
// name.
func (node *PlaceNode) Name() string {
	name, _, _, _ := node.JurisdictionalEntities()

	return name
}

// County is the second part of the JurisdictionalName().
//
// County will only return a non-empty response if the JurisdictionalName is
// exactly in the form of "Name,County,State,Country".
func (node *PlaceNode) County() string {
	_, county, _, _ := node.JurisdictionalEntities()

	return county
}

// State is the third part of the JurisdictionalName().
//
// State will only return a non-empty response if the JurisdictionalName is
// exactly in the form of "Name,County,State,Country".
func (node *PlaceNode) State() string {
	_, _, state, _ := node.JurisdictionalEntities()

	return state
}

// Country is the forth part of the JurisdictionalName().
//
// Country will only return a non-empty response if the JurisdictionalName is
// exactly in the form of "Name,County,State,Country" or the country can be
// identified from the list of Countries.
func (node *PlaceNode) Country() string {
	_, _, _, country := node.JurisdictionalEntities()

	if country != "" {
		return country
	}

	// If the country is empty it is likely because the place is not formatted
	// into four jurisdictional entities. In this case we will try to find the
	// country by looking at the suffix of the place name.
	nameWithoutPunctuation := strings.Trim(node.JurisdictionalName(), ",. ")
	lowerCaseName := strings.ToLower(nameWithoutPunctuation)

	for _, c := range Countries {
		if strings.HasSuffix(lowerCaseName, strings.ToLower(c)) {
			return c
		}
	}

	return ""
}

// Format shows the jurisdictional entities that are named in a sequence from
// the lowest to the highest jurisdiction.
//
// The jurisdictions are separated by commas, and any jurisdiction's name that
// is missing is still accounted for by a comma.
//
// When a PLAC.FORM structure is included in the HEADER of a GEDCOM
// transmission, it implies that all place names follow this jurisdictional
// format and each jurisdiction is accounted for by a comma, whether the name is
// known or not.
//
// When the PLAC.FORM is subordinate to an event, it temporarily overrides the
// implications made by the PLAC.FORM structure stated in the HEADER. This usage
// is not common and, therefore, not encouraged. It should only be used when a
// system has over-structured its place-names.
//
// See JurisdictionalName() for a more reliable way to determine the Format.
func (node *PlaceNode) Format() *FormatNode {
	n := First(NodesWithTag(node, TagFormat))

	if IsNil(n) {
		return nil
	}

	return n.(*FormatNode)
}

// PhoneticVariations of the place name are written in the same form as was the
// place name used in the superior PlaceNode value, but phonetically written
// using the method indicated by the subordinate Type.
//
// For example if hiragana was used to provide a reading of a name written in
// kanji, then the Type value would indicate kana.
//
// See PhoneticVariationTypeHangul and PhoneticVariationTypeKana.
func (node *PlaceNode) PhoneticVariations() []*PhoneticVariationNode {
	t := (*PhoneticVariationNode)(nil)

	return castNodesWithTag(node, TagPhonetic, t).([]*PhoneticVariationNode)
}

// RomanizedVariations of the place name are written in the same form prescribed
// for the place name used in the superior PlaceNode context.
//
// The method used to romanize the name is indicated by the Value of the
// subordinate Type().
//
// For example if romaji was used to provide a reading of a place name written
// in kanji, then the Type subordinate to the RomanizedVariationNode would
// indicate "romaji".
func (node *PlaceNode) RomanizedVariations() []*RomanizedVariationNode {
	t := (*RomanizedVariationNode)(nil)

	return castNodesWithTag(node, TagRomanized, t).([]*RomanizedVariationNode)
}

func (node *PlaceNode) Map() *MapNode {
	n := First(NodesWithTag(node, TagMap))

	if IsNil(n) {
		return nil
	}

	return n.(*MapNode)
}

func (node *PlaceNode) Notes() []*NoteNode {
	t := (*NoteNode)(nil)

	return castNodesWithTag(node, TagNote, t).([]*NoteNode)
}

// JurisdictionalEntities returns the name, county, state and country.
//
// See JurisdictionalName() for a full explanation.
func (node *PlaceNode) JurisdictionalEntities() (string, string, string, string) {
	jurisdictionalName := node.JurisdictionalName()
	placeParts := strings.Split(jurisdictionalName, ",")

	if len(placeParts) != 4 {
		placeParts = []string{jurisdictionalName, "", "", ""}
	}

	name := strings.TrimSpace(placeParts[0])
	county := strings.TrimSpace(placeParts[1])
	state := strings.TrimSpace(placeParts[2])
	country := strings.TrimSpace(placeParts[3])

	return name, county, state, country
}
