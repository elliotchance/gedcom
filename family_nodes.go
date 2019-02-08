package gedcom

type FamilyNodes []*FamilyNode

func (nodes FamilyNodes) ByPointer(pointer string) *FamilyNode {
	for _, node := range nodes {
		if node.Pointer() == pointer {
			return node
		}
	}

	return nil
}
