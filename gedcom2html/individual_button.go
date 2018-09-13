package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// individualButton is a large coloured button that links to an individuals
// page. It contains the same and some date information. This is also used to
// represent unknown or missing individuals.
type individualButton struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func newIndividualButton(document *gedcom.Document, individual *gedcom.IndividualNode) *individualButton {
	return &individualButton{
		individual: individual,
		document:   document,
	}
}

func (c *individualButton) String() string {
	name := html.NewIndividualName(c.individual, false, html.UnknownEmphasis).String()

	onclick := ""
	if c.individual != nil {
		onclick = html.Sprintf(`onclick="location.href='%s'"`,
			pageIndividual(c.document, c.individual))
	}

	eventDates := html.NewIndividualDates(c.individual, false)

	// If the individual is living we need to hide all their information.
	if c.individual != nil && c.individual.IsLiving() {
		name = "<em>Hidden</em>"
		onclick = ""
	}

	return html.Sprintf(`
		<button type="button" class="btn btn-outline-%s btn-block" %s>
			<strong>%s</strong><br/>
			%s&nbsp;
		</button>
	`, colorClassForIndividual(c.individual), onclick, name, eventDates)
}
