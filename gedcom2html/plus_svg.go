package main

import "fmt"

// plusSVG draws a "+" as an SVG with each line of the "+" being optional.
type plusSVG struct {
	top, left, right, bottom bool
}

func newPlusSVG(top, left, right, bottom bool) *plusSVG {
	return &plusSVG{
		top:    top,
		left:   left,
		right:  right,
		bottom: bottom,
	}
}

func (c *plusSVG) String() string {
	// The "+" is constructed of two lines. Each of the lines need 4 coordinates
	// to represent the start and end points of X and Y. The values represent
	// percentages.
	//
	// By default the lines are set to render the none of the sides of the "+".
	// So it would render a small square dot in the middle. Each of the
	// activated options will adjust the appropriate sides to extend the lines
	// in that direction.

	hLineX1, hLineX2, hLineY1, hLineY2 := 50, 50, 50, 50
	vLineX1, vLineX2, vLineY1, vLineY2 := 50, 50, 50, 50

	if c.top {
		vLineY1 = 20
	}

	if c.left {
		hLineX1 = 0
	}

	if c.right {
		hLineX2 = 100
	}

	if c.bottom {
		vLineY2 = 80
	}

	return fmt.Sprintf(`
		<svg style="width: 100%%; height: 75px">
			<line x1="%d%%" y1="%d%%" x2="%d%%" y2="%d%%" style="stroke:rgb(0,0,0);stroke-width:3" />
			<line x1="%d%%" y1="%d%%" x2="%d%%" y2="%d%%" style="stroke:rgb(0,0,0);stroke-width:3" />
		</svg>
	`, hLineX1, hLineY1, hLineX2, hLineY2, vLineX1, vLineY1, vLineX2, vLineY2)
}
