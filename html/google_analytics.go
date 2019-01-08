package html

import "io"

type GoogleAnalytics struct {
	id string
}

func NewGoogleAnalytics(id string) *GoogleAnalytics {
	return &GoogleAnalytics{
		id: id,
	}
}

func (c *GoogleAnalytics) WriteTo(w io.Writer) (int64, error) {
	if c.id == "" {
		return writeNothing()
	}

	return writeSprintf(w, `<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=%s"></script>
<script>
window.dataLayer = window.dataLayer || [];
function gtag(){dataLayer.push(arguments);}
gtag('js', new Date());

gtag('config', '%s');
</script>`, c.id, c.id)
}
