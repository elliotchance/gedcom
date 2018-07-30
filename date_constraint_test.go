package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
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
