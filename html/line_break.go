package html

type LineBreak struct{}

func NewLineBreak() *LineBreak {
	return &LineBreak{}
}

func (c *LineBreak) String() string {
	return `<br/>`
}
