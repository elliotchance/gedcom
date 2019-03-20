package core

type FileWriter interface {
	WriteFile(file *File) error
}
