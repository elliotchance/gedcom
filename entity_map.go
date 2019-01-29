package gedcom

type entityMap map[interface{}]interface{}

func (m entityMap) GetOrAssign(n interface{}, assign func() interface{}) interface{} {
	if existing, ok := m[n]; ok {
		return existing
	}

	newNode := assign()
	m[n] = newNode

	return newNode
}
