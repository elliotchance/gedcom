package gedcom

type Warnings []Warning

func (ws Warnings) Strings() (ss []string) {
	for _, w := range ws {
		ss = append(ss, w.String())
	}

	return
}

func (ws Warnings) MarshalQ() interface{} {
	out := []interface{}{}
	for _, warning := range ws {
		out = append(out, warning.MarshalQ())
	}

	return out
}
