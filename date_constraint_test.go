package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/stretchr/testify/assert"
)

func TestDateConstraint_String(t *testing.T) {
	tests := map[gedcom.DateConstraint]string{
		gedcom.DateConstraintExact:  "",
		gedcom.DateConstraintAbout:  "Abt.",
		gedcom.DateConstraintAfter:  "Aft.",
		gedcom.DateConstraintBefore: "Bef.",
		gedcom.DateConstraint(35):   "",
	}

	for constraint, expected := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equalf(t, expected, constraint.String(), "%#+v", constraint)
		})
	}
}
