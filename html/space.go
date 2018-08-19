package html

// Space is an empty row used as a white space separator between other page
// rows.
type Space struct{}

func NewSpace() *Space {
	return &Space{}
}

func (c *Space) String() string {
	return `
	<div class="row">
        <div class="col">
            &nbsp;
        </div>
    </div>`
}
