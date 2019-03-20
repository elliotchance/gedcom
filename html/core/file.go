package core

type File struct {
	Name      string
	Component Component
}

func NewFile(name string, component Component) *File {
	return &File{
		Name:      name,
		Component: component,
	}
}
