package html

type LivingVisibility string

const (
	LivingVisibilityShow        = "show"
	LivingVisibilityHide        = "hide"
	LivingVisibilityPlaceholder = "placeholder"
)

func NewLivingVisibility(lv string) LivingVisibility {
	switch LivingVisibility(lv) {
	case LivingVisibilityShow,
		LivingVisibilityHide,
		LivingVisibilityPlaceholder:
		return LivingVisibility(lv)
	}

	panic("invalid LivingVisibility: " + lv)
}
