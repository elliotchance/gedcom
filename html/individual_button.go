package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

// IndividualButton is a large coloured button that links to an individuals
// page. It contains the same and some date information. This is also used to
// represent unknown or missing individuals.
type IndividualButton struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
	visibility LivingVisibility
}

func NewIndividualButton(document *gedcom.Document, individual *gedcom.IndividualNode, visibility LivingVisibility) *IndividualButton {
	return &IndividualButton{
		individual: individual,
		document:   document,
		visibility: visibility,
	}
}

func (c *IndividualButton) WriteHTMLTo(w io.Writer) (int64, error) {
	if c.individual.IsLiving() {
		switch c.visibility {
		case LivingVisibilityHide:
			return writeNothing()

		case LivingVisibilityShow, LivingVisibilityPlaceholder:
			// Proceed.
		}
	}

	var name core.Component = NewIndividualName(c.individual, c.visibility, UnknownEmphasis)

	onclick := ""
	if c.individual != nil {
		onclick = fmt.Sprintf(`location.href='%s'`,
			PageIndividual(c.document, c.individual, c.visibility))
	}

	eventDates := NewIndividualDates(c.individual, c.visibility)

	isLiving := c.individual != nil && c.individual.IsLiving()
	if isLiving {
		switch c.visibility {
		case LivingVisibilityHide:
			return writeNothing()

		case LivingVisibilityShow:
			// Proceed.

		case LivingVisibilityPlaceholder:
			name = core.NewHTML("<em>Hidden</em>")
			onclick = ""
		}
	}

	return core.NewTag("button", map[string]string{
		"class":   fmt.Sprintf("btn btn-outline-%s btn-block", colorClassForIndividual(c.individual)),
		"type":    "button",
		"onclick": onclick,
	}, core.NewComponents(
		core.NewTag("strong", nil, name),
		core.NewLineBreak(),
		eventDates,
		core.NewEmpty(),
	)).WriteHTMLTo(w)
}
