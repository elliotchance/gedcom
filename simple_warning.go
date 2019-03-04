package gedcom

type SimpleWarning struct {
	context WarningContext
}

func (w *SimpleWarning) SetContext(context WarningContext) {
	w.context = context
}

func (w *SimpleWarning) Context() WarningContext {
	return w.context
}
