package gedcom

type Warnings []Warning

func (ws Warnings) Strings() (ss []string) {
	for _, w := range ws {
		ss = append(ss, w.String())
	}

	return
}
