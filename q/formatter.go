package q

// Formatter is used to write the result to stream.
type Formatter interface {
	Write(result interface{}) error
}
