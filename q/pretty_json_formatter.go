package q

import (
	"encoding/json"
	"io"
)

type PrettyJSONFormatter struct {
	Writer io.Writer
}

func (f *PrettyJSONFormatter) Write(result interface{}) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Writer.Write(append(data, '\n'))

	return err
}
