package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// parentButtons show two buttons separated by a "T" to be placed above the
// large individuals name.
type parentButtons struct {
	family   *gedcom.FamilyNode
	document *gedcom.Document
}

func newParentButtons(document *gedcom.Document, family *gedcom.FamilyNode) *parentButtons {
	return &parentButtons{
		family:   family,
		document: document,
	}
}

func (c *parentButtons) String() string {
	husband := newIndividualButton(c.document, c.family.Husband())
	wife := newIndividualButton(c.document, c.family.Wife())
	svg := newPlusSVG(false, true, true, true)
	space := html.NewSpace()

	return html.Sprintf(`
		<div class="row">
		   <div class="col-5">
		       %s
		   </div>
		   <div class="col-2">
               %s
		   </div>
		   <div class="col-5">
		       %s
		   </div>
		</div>
		%s`,
		husband, svg, wife, space)
}
