package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYears(t *testing.T) {
	tests := []struct {
		v    interface{}
		want float64
	}{
		{nil, 0},
		{gedcom.Yearer(nil), 0},
		{gedcom.Date{1, 1, 1789, false, gedcom.DateConstraintExact, nil}, 1789.0027322404371},
		{gedcom.NewDateNode("Foo"), 0},
		{gedcom.NewDateNode("31 Jan 1435"), 1435.0846994535518},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Years(test.v))
		})
	}
}
