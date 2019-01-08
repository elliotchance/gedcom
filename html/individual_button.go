package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"io"
)

// IndividualButton is a large coloured button that links to an individuals
// page. It contains the same and some date information. This is also used to
// represent unknown or missing individuals.
type IndividualButton struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func NewIndividualButton(document *gedcom.Document, individual *gedcom.IndividualNode) *IndividualButton {
	return &IndividualButton{
		individual: individual,
		document:   document,
	}
}

func (c *IndividualButton) WriteTo(w io.Writer) (int64, error) {
	var name Component = NewIndividualName(c.individual, false, UnknownEmphasis)

	onclick := ""
	if c.individual != nil {
		onclick = fmt.Sprintf(`location.href='%s'`,
			PageIndividual(c.document, c.individual))
	}

	eventDates := NewIndividualDates(c.individual, false)

	// If the individual is living we need to hide all their information.
	if c.individual != nil && c.individual.IsLiving() {
		name = NewHTML("<em>Hidden</em>")
		onclick = ""
	}

	return NewTag("button", map[string]string{
		"class":   fmt.Sprintf("btn btn-outline-%s btn-block", colorClassForIndividual(c.individual)),
		"type":    "button",
		"onclick": onclick,
	}, NewComponents(
		NewTag("strong", nil, name),
		NewLineBreak(),
		eventDates,
		NewEmpty(),
	)).WriteTo(w)
}
