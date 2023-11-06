package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
)

func TestSourceNode_Title(t *testing.T) {
	Title := tf.Function(t, (*gedcom.SourceNode).Title)

	Title((*gedcom.SourceNode)(nil)).Returns("")
}
