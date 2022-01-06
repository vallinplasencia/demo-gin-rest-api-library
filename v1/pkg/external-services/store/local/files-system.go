package local

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	apstoreabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/store/abstract"
)

// FilesSytem permite el acceso al sistema de archivos localmente.
type FilesSytem struct {
	// Directory directorio principal donde se guardarn los archivos subidos
	Directory string
}

// New ...
func New(c *Config) (apstoreabstract.Store, error) {
	return &FilesSytem{
		Directory: strings.TrimSuffix(c.DestinationTarget, string(os.PathSeparator)),
	}, nil
}

// Save guarda un archivo en el sistema de ficheros localmente.
func (u *FilesSytem) Save(pathWithName string, r io.Reader) (*apstoreabstract.FileData, error) {
	osSeparator := string(os.PathSeparator)
	if !strings.HasPrefix(pathWithName, osSeparator) {
		pathWithName = fmt.Sprintf("%s%s", osSeparator, pathWithName)
	}
	pathDir := fmt.Sprintf("%s%s", u.Directory, pathWithName[:strings.LastIndex(pathWithName, osSeparator)])

	if _, e := os.Stat(pathDir); e != nil {
		if os.IsNotExist(e) {
			e := os.MkdirAll(pathDir, 0755)
			if e != nil {
				return nil, e
			}
		} else {
			return nil, e
		}
	}
	dst, e := os.Create(fmt.Sprintf("%s%s", u.Directory, pathWithName))

	if e != nil {
		return nil, e
	}
	defer dst.Close()

	if _, e = io.Copy(dst, r); e != nil {
		os.Remove(pathWithName) // elimino el archivo creado con anterioridad como posible destino
		return nil, e
	}
	return &apstoreabstract.FileData{
		Path:       pathWithName,
		UploadedAt: time.Now().UTC().Unix(),
	}, nil
}

// Remove elimina un archivo en el sistema de ficheros localmente donde path es
// la ruta completa del archivo a eliminar
func (u *FilesSytem) Remove(path string) error {
	return os.Remove(fmt.Sprintf("%s%s", u.Directory, path))
}
