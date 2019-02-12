package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
	"time"
)

func TestDuration_String(t *testing.T) {
	String := tf.NamedFunction(t, "Duration_String", gedcom.Duration.String)

	oneDay := 24 * time.Hour

	String(0).Returns("one day")
	String(time.Minute).Returns("one day")
	String(time.Hour * 30).Returns("2 days")
	String(oneDay * 8).Returns("8 days")
	String(oneDay * 21).Returns("21 days")
	String(oneDay * 30).Returns("30 days")
	String(oneDay * 31).Returns("one month and one day")
	String(oneDay * 32).Returns("one month and 2 days")
	String(oneDay * 60).Returns("one month and 30 days")
	String(oneDay * 61).Returns("2 months and one day")
	String(oneDay * 70).Returns("2 months and 10 days")
	String(oneDay * 360).Returns("11 months and 26 days")
	String(oneDay * 365).Returns("one year")
	String(oneDay * 366).Returns("one year and one day")
	String(oneDay * 400).Returns("one year and one month and 5 days")
	String(oneDay * 440).Returns("one year and 2 months and 15 days")
	String(oneDay * 1000).Returns("2 years and 8 months and 27 days")
}
