package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

// ParentButtons show two buttons separated by a "T" to be placed above the
// large individuals name.
type ParentButtons struct {
	family     *gedcom.FamilyNode
	document   *gedcom.Document
	visibility LivingVisibility
	placesMap  map[string]*place
}

func NewParentButtons(document *gedcom.Document, family *gedcom.FamilyNode, visibility LivingVisibility, placesMap map[string]*place) *ParentButtons {
	return &ParentButtons{
		family:     family,
		document:   document,
		visibility: visibility,
		placesMap:  placesMap,
	}
}

func (c *ParentButtons) WriteHTMLTo(w io.Writer) (int64, error) {
	husband := NewIndividualButton(c.document, c.family.Husband().Individual(),
		c.visibility, c.placesMap)
	wife := NewIndividualButton(c.document, c.family.Wife().Individual(),
		c.visibility, c.placesMap)
	svg := NewPlusSVG(false, true, true, true)
	space := core.NewSpace()

	return core.NewComponents(
		core.NewRow(
			core.NewColumn(5, husband),
			core.NewColumn(2, svg),
			core.NewColumn(5, wife),
		),
		space,
	).WriteHTMLTo(w)
}
