package gedcom

type QMarshaller interface {
	MarshalQ() interface{}
}

func MarshalQ(in interface{}) interface{} {
	if m, ok := in.(QMarshaller); ok {
		return m.MarshalQ()
	}

	return in
}
