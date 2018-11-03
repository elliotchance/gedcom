package html

type HTML struct {
	s string
}

func NewHTML(s string) *HTML {
	return &HTML{
		s: s,
	}
}

func (c *HTML) String() string {
	return c.s
}
