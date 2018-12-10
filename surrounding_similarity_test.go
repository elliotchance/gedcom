package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestSurroundingSimilarity_WeightedSimilarity(t *testing.T) {
	WS := tf.Function(t, (*gedcom.SurroundingSimilarity).WeightedSimilarity)

	WS(gedcom.NewSurroundingSimilarity(0.0, 0.0, 0.0, 0.0)).Returns(0.0)
	WS(gedcom.NewSurroundingSimilarity(1.0, 0.0, 0.0, 0.0)).Returns(0.06666666666666666)
	WS(gedcom.NewSurroundingSimilarity(0.0, 1.0, 0.0, 0.0)).Returns(0.8)
	WS(gedcom.NewSurroundingSimilarity(0.0, 0.0, 1.0, 0.0)).Returns(0.06666666666666666)
	WS(gedcom.NewSurroundingSimilarity(0.0, 0.0, 0.0, 1.0)).Returns(0.06666666666666666)
	WS(gedcom.NewSurroundingSimilarity(1.0, 0.5, 1.0, 1.0)).Returns(0.6)
	WS(gedcom.NewSurroundingSimilarity(0.8, 0.8, 0.8, 0.8)).Returns(0.8000000000000002)
	WS(gedcom.NewSurroundingSimilarity(1.0, 1.0, 1.0, 1.0)).Returns(1.0)
}
