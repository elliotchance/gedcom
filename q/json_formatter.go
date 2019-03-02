package q

import (
	"encoding/json"
	"io"
)

type JSONFormatter struct {
	Writer io.Writer
}

func (f *JSONFormatter) Write(result interface{}) error {
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	_, err = f.Writer.Write(append(data, '\n'))

	return err
}
