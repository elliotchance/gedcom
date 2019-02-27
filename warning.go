package gedcom

import (
	"fmt"
)

type Warning interface {
	fmt.Stringer
	QMarshaller

	Name() string
	SetContext(context WarningContext)
}
