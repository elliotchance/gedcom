package main

import (
	"github.com/elliotchance/tf"
	"testing"
)

func Test_badgePill_String(t *testing.T) {
	String := tf.Function(t, (*badgePill).String)

	String(newBadgePill("green", "myclass", "123")).
		Returns(`<span class="badge badge-pill badge-green myclass">123</span>`)
}
