package gedcom

type IndividualNodePairs []*IndividualNodePair

func (pairs IndividualNodePairs) Has(findPair *IndividualNodePair) bool {
	for _, pair := range pairs {
		if pair.Left.Is(findPair.Left) && pair.Right.Is(findPair.Right) {
			return true
		}

		if pair.Left.Is(findPair.Right) && pair.Right.Is(findPair.Left) {
			return true
		}
	}

	return false
}
