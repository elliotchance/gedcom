package html_test

import (
	"bytes"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIndividualEvents(t *testing.T) {
	doc, err := gedcom.NewDocumentFromString(`
0 @I492@ INDI
1 NAME Eva Ellen /Preece/
2 SOUR @S14@
3 DATA
4 TEXT Record for Eva Ellen Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=FreeBMDBirth&h=28959069&indiv=try
2 SOUR @S9@
3 PAGE Class: RG13; Piece: 2422; Folio: 119; Page: 10
3 DATA
4 TEXT Record for John Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=uki1901&h=13142652&indiv=try
2 SOUR @S1@
3 PAGE Database online.
3 DATA
4 TEXT Record for Eliza Farley
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=pubmembertrees&h=-442570170&indiv=try
2 SOUR @S3@
3 PAGE Class: RG14; Piece: 15243; Schedule Number: 109
3 DATA
4 TEXT Record for Eva Ellen Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=1911England&h=56002892&indiv=try
2 SOUR @S29@
3 PAGE Gloucestershire Archives; Gloucester, England; Reference Numbers: P265 IN 1/7
3 DATA
4 TEXT Record for Eva Ellen Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=GloucBapt&h=632299&indiv=try
1 SEX F
1 NAME Dorothy Ellen
2 TYPE aka
1 RESI Age: 3; Relation to Head of House: Daughter
2 DATE 1901
2 PLAC Redmarley D'Abitot, Worcestershire, England
3 MAP
4 LATI N51.98
4 LONG W2.3615
2 SOUR @S9@
3 PAGE Class: RG13; Piece: 2422; Folio: 119; Page: 10
3 DATA
4 TEXT Record for John Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=uki1901&h=13142652&indiv=try
1 BIRT
2 DATE 1 FEB 1898
2 PLAC Redmarley D'Abitot, Worcestershire, England
3 MAP
4 LATI N51.98
4 LONG W2.3615
2 SOUR @S14@
3 DATA
4 TEXT Record for Eva Ellen Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=FreeBMDBirth&h=28959069&indiv=try
2 SOUR @S1@
3 PAGE Database online.
3 DATA
4 TEXT Record for Eliza Farley
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=pubmembertrees&h=-442570170&indiv=try
2 SOUR @S9@
3 PAGE Class: RG13; Piece: 2422; Folio: 119; Page: 10
3 DATA
4 TEXT Record for John Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=uki1901&h=13142652&indiv=try
2 SOUR @S3@
3 PAGE Class: RG14; Piece: 15243; Schedule Number: 109
3 DATA
4 TEXT Record for Eva Ellen Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=1911England&h=56002892&indiv=try
1 BAPM
2 DATE 5 JUN 1898
2 PLAC Redmarley D'Abitot, Worcestershire, England
3 MAP
4 LATI N51.98
4 LONG W2.3615
2 SOUR @S29@
3 PAGE Gloucestershire Archives; Gloucester, England; Reference Numbers: P265 IN 1/7
3 DATA
4 TEXT Record for Eva Ellen Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=GloucBapt&h=632299&indiv=try
1 DEAT alive in 1960
1 RESI Age: 13; Relation to Head of House: Niece
2 DATE 2 APR 1911
2 PLAC Newent, Gloucestershire, England
3 MAP
4 LATI N51.9309
4 LONG W2.4054
2 SOUR @S3@
3 PAGE Class: RG14; Piece: 15243; Schedule Number: 109
3 DATA
4 TEXT Record for Eva Ellen Preece
3 _LINK http://search.ancestry.co.uk/cgi-bin/sse.dll?db=1911England&h=56002892&indiv=try
1 FAMC @F138@`)
	require.NoError(t, err)

	component := html.NewIndividualEvents(doc, doc.Individuals()[0],
		html.LivingVisibilityPlaceholder, nil)
	buf := bytes.NewBuffer(nil)
	_, err = component.WriteHTMLTo(buf)
	require.NoError(t, err)

	assertTextByXPath(t, buf.String(), "//table//text()", []string{
		"Age", "Type", "Date", "Place", "Description",
		"0y", "Birth", "1 Feb 1898", "Redmarley D'Abitot,  Worcestershire,  England", "\u00a0",
		"~ 2y 10m", "Residence", "1901", "Redmarley D'Abitot,  Worcestershire,  England", "\u00a0",
		"~ 0y 4m", "Baptism", "5 Jun 1898", "Redmarley D'Abitot,  Worcestershire,  England", "\u00a0",
		"~ 13y 1m", "Residence", "2 Apr 1911", "Newent,  Gloucestershire,  England", "\u00a0",
		"Death", "\u00a0",
	})
}
