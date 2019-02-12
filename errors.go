package gedcom

import (
	"errors"
	"fmt"
	"strings"
)

type Errors []error

func NewErrors(errors ...error) (errs Errors) {
	for _, err := range errors {
		if err != nil {
			errs = append(errs, err)
		}
	}

	return
}

func (e Errors) Errors() (errors []string) {
	for _, err := range e {
		errors = append(errors, err.Error())
	}

	return
}

func (e Errors) Error() string {
	if len(e) == 1 {
		return e[0].Error()
	}

	errs := strings.Join(e.Errors(), "; ")

	return fmt.Sprintf("%d errors: %s", len(e), errs)
}

func (e Errors) Err() error {
	if len(e) == 0 {
		return nil
	}

	return errors.New(e.Error())
}
