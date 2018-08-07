package gedcom

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestSurroundingSimilarity_WeightedSimilarity(t *testing.T) {
	WS := tf.Function(t, SurroundingSimilarity.WeightedSimilarity)

	WS(SurroundingSimilarity{}).Returns(0.0)
	WS(SurroundingSimilarity{0.0, 0.0, 0.0, 0.0}).Returns(0.0)
	WS(SurroundingSimilarity{1.0, 0.0, 0.0, 0.0}).Returns(0.16666666666666666)
	WS(SurroundingSimilarity{0.0, 1.0, 0.0, 0.0}).Returns(0.5)
	WS(SurroundingSimilarity{0.0, 0.0, 1.0, 0.0}).Returns(0.16666666666666666)
	WS(SurroundingSimilarity{0.0, 0.0, 0.0, 1.0}).Returns(0.16666666666666666)
	WS(SurroundingSimilarity{1.0, 0.5, 1.0, 1.0}).Returns(0.75)
	WS(SurroundingSimilarity{0.8, 0.8, 0.8, 0.8}).Returns(0.8)
	WS(SurroundingSimilarity{1.0, 1.0, 1.0, 1.0}).Returns(1.0)
}
