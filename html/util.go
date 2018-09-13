package html

import (
	"fmt"
)

func Sprintf(format string, args ...interface{}) string {
	newArgs := make([]interface{}, len(args))
	for i, arg := range args {
		if a, ok := arg.(fmt.Stringer); ok {
			newArgs[i] = a.String()
		} else {
			newArgs[i] = arg
		}
	}

	return fmt.Sprintf(format, newArgs...)
}
