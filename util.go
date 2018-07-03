package gedcom

func valueToPointer(val string) string {
	if len(val) > 2 && val[0] == '@' && val[len(val)-1] == '@' {
		return val[1 : len(val)-1]
	}

	return ""
}
