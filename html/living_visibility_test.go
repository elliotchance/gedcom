package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLivingVisibility(t *testing.T) {
	NewLivingVisibility := tf.Function(t, html.NewLivingVisibility)

	NewLivingVisibility("show").Returns(html.LivingVisibilityShow)
	NewLivingVisibility("hide").Returns(html.LivingVisibilityHide)
	NewLivingVisibility("placeholder").Returns(html.LivingVisibilityPlaceholder)

	assert.PanicsWithValue(t, "invalid LivingVisibility: foo", func() {
		html.NewLivingVisibility("foo")
	})
}
