package abstract

import (
	"io"
)

// Store define los metodos q maneja los archivos
type Store interface {
	Save(pathWithName string, reader io.Reader) (*FileData, error)
	Remove(path string) error
}

// FileData datos del archivo que se subio
type FileData struct {
	// Path ruta interna del archivo q se subio
	Path string
	// UploadedAt fecha unix en la cual se termino de subir el archivo
	UploadedAt int64
}
