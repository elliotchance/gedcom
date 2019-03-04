package gedcom

import (
	"fmt"
)

type Warning interface {
	fmt.Stringer

	Name() string
	SetContext(context WarningContext)
	Context() WarningContext
}
