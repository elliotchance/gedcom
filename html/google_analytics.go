package html

import (
	"fmt"
)

type GoogleAnalytics struct {
	id string
}

func newGoogleAnalytics(id string) *GoogleAnalytics {
	return &GoogleAnalytics{
		id: id,
	}
}

func (c *GoogleAnalytics) String() string {
	if c.id == "" {
		return ""
	}

	return fmt.Sprintf(`<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=%s"></script>
<script>
window.dataLayer = window.dataLayer || [];
function gtag(){dataLayer.push(arguments);}
gtag('js', new Date());

gtag('config', '%s');
</script>`, c.id, c.id)
}
