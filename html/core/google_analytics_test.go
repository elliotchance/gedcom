package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestGoogleAnalytics_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "GoogleAnalytics")

	c(core.NewGoogleAnalytics("the-id")).
		Returns("<!-- Global site tag (gtag.js) - Google Analytics -->\n<script async src=\"https://www.googletagmanager.com/gtag/js?id=the-id\"></script>\n<script>\nwindow.dataLayer = window.dataLayer || [];\nfunction gtag(){dataLayer.push(arguments);}\ngtag('js', new Date());\n\ngtag('config', 'the-id');\n</script>")
}
