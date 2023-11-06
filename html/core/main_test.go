package core_test

import (
	"bytes"
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func testComponent(t *testing.T, name string) func(args ...interface{}) *tf.F {
	return tf.NamedFunction(t, name+"_WriteHTMLTo", func(c core.Component) string {
		buf := bytes.NewBuffer(nil)
		n, err := c.WriteHTMLTo(buf)
		assert.NoError(t, err)

		data := buf.Bytes()
		assert.Equal(t, int64(len(data)), n)

		return string(data)
	})
}
