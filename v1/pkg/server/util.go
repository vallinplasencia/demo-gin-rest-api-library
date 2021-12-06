package server

// StoreFileType donde se almacenan los archivos subidos
type StoreFileType string

const (
	// StoreFilesSystemLocal en el sistema de ficheros de la pc donde se ejecuta la app
	StoreFilesSystemLocal StoreFileType = "files-system-local"
	// StoreAwsS3 en un bucket de s3 en AWS
	StoreAwsS3 StoreFileType = "aws-s3"
)
