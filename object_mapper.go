package gedcom

// ObjectMapper allows objects to be represented a map. This is important when
// serializing and examining objects.
type ObjectMapper interface {
	ObjectMap() map[string]interface{}
}
