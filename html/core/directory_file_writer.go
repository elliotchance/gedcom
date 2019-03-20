package core

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

type DirectoryFileWriter struct {
	outputDir     string
	WillWriteFile func(file *File)
}

func NewDirectoryFileWriter(outputDir string) *DirectoryFileWriter {
	return &DirectoryFileWriter{
		outputDir: outputDir,
	}
}

func (writer *DirectoryFileWriter) WriteFile(file *File) error {
	path := fmt.Sprintf("%s/%s", writer.outputDir, file.Name)

	if writer.WillWriteFile != nil {
		writer.WillWriteFile(file)
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}

	file.Component.WriteHTMLTo(out)

	return out.Close()
}

func (writer *DirectoryFileWriter) fileSha1(path string) (string, error) {
	f, err := os.Open(writer.outputDir + "/" + path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
