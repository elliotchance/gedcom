package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterFlags_Filter(t *testing.T) {
	t.Run("NoDuplicateNames", func(t *testing.T) {
		doc1 := gedcom.NewDocument()
		doc1.AddIndividual("P1",
			gedcom.NewNameNode("Bob /Smith/"),
			gedcom.NewNameNode("Jane /Smith/"),
			gedcom.NewNameNode("Bob /Smith/"))

		doc2 := gedcom.NewDocument()
		doc2.AddIndividual("P1",
			gedcom.NewNameNode("Bob /Smith/"),
			gedcom.NewNameNode("Jane /Smith/"))

		ff := &gedcom.FilterFlags{
			NoDuplicateNames: true,
			NameFormat:       "unmodified",
		}

		assert.Equal(t, doc2.Nodes()[0].GEDCOMString(0),
			ff.Filter(doc1.Nodes()[0]).GEDCOMString(0))
	})

	t.Run("NoDuplicateNamesWithModifiedNames", func(t *testing.T) {
		doc1 := gedcom.NewDocument()
		doc1.AddIndividual("P1",
			gedcom.NewNameNode("Bob /Smith/"),
			gedcom.NewNameNode("Jane /Smith/"),
			gedcom.NewNameNode("Bob /Smith/"))

		doc2 := gedcom.NewDocument()
		doc2.AddIndividual("P1",
			gedcom.NewNameNode("Bob /Smith/"),
			gedcom.NewNameNode("Jane /Smith/"))

		ff := &gedcom.FilterFlags{
			NoDuplicateNames: true,
			NameFormat:       string(gedcom.NameFormatGEDCOM),
		}

		assert.Equal(t, doc2.Nodes()[0].GEDCOMString(0),
			ff.Filter(doc1.Nodes()[0]).GEDCOMString(0))
	})
}
