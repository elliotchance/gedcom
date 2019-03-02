package gedcom

type SimpleWarning struct {
	Context WarningContext
}

func (w *SimpleWarning) SetContext(context WarningContext) {
	w.Context = context
}
