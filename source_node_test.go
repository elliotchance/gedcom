package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestSourceNode_Title(t *testing.T) {
	Title := tf.Function(t, (*gedcom.SourceNode).Title)

	Title((*gedcom.SourceNode)(nil)).Returns("")
}
