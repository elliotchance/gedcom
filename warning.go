package gedcom

import "fmt"

type Warning interface {
	fmt.Stringer
	Name() string
}
