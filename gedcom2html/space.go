package main

// space is an empty row used as a white space separator between other page
// rows.
type space struct{}

func newSpace() *space {
	return &space{}
}

func (c *space) String() string {
	return `
	<div class="row">
        <div class="col">
            &nbsp;
        </div>
    </div>`
}
