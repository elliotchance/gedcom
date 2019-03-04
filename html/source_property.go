package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type SourceProperty struct {
	document *gedcom.Document
	node     gedcom.Node
}

func NewSourceProperty(document *gedcom.Document, node gedcom.Node) *SourceProperty {
	return &SourceProperty{
		document: document,
		node:     node,
	}
}

func (c *SourceProperty) WriteHTMLTo(w io.Writer) (int64, error) {
	tag := c.node.Tag().String()
	value := c.node.Value()

	components := []core.Component{
		core.NewTableHead(tag, value),
	}

	for _, node := range c.node.Nodes() {
		components = append(components, NewSourceProperty(c.document, node))
	}

	return core.NewComponents(components...).WriteHTMLTo(w)
}
