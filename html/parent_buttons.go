package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

// ParentButtons show two buttons separated by a "T" to be placed above the
// large individuals name.
type ParentButtons struct {
	family   *gedcom.FamilyNode
	document *gedcom.Document
}

func NewParentButtons(document *gedcom.Document, family *gedcom.FamilyNode) *ParentButtons {
	return &ParentButtons{
		family:   family,
		document: document,
	}
}

func (c *ParentButtons) WriteTo(w io.Writer) (int64, error) {
	husband := NewIndividualButton(c.document, c.family.Husband())
	wife := NewIndividualButton(c.document, c.family.Wife())
	svg := NewPlusSVG(false, true, true, true)
	space := NewSpace()

	return NewComponents(
		NewRow(
			NewColumn(5, husband),
			NewColumn(2, svg),
			NewColumn(5, wife),
		),
		space,
	).WriteTo(w)
}
