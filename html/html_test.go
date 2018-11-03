package html

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestHTML_String(t *testing.T) {
	String := tf.Function(t, (*HTML).String)

	String(NewHTML("foo")).Returns(`foo`)
	String(NewHTML(`"Fran & Freddie's Diner" <tasty@example.com>`)).
		Returns(`"Fran & Freddie's Diner" <tasty@example.com>`)
}
