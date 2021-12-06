package s3

import (
	"context"
	"errors"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	apstoreabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/store/abstract"
)

// AwsS3 permite el acceso a los bucket de S3
type AwsS3 struct {
	// bucket el nombre del bucket a subir los archivos
	bucket string
	config *aws.Config
}

// New retorna intancia q implementa apstore.Store.
func New(c *config) (apstoreabstract.Store, error) {
	if len(c.DestinationTarget) == 0 {
		return nil, errors.New("destination target is empty")
	}
	// load data for auth from ...
	cfg, e := awsconfig.LoadDefaultConfig(context.TODO())
	if e != nil {
		return nil, e
	}
	return &AwsS3{
		bucket: c.DestinationTarget,
		config: &cfg,
	}, nil
}

// Save guarda un archivo en el bucket.
func (u *AwsS3) Save(pathWithName string, r io.Reader) (*apstoreabstract.FileData, error) {
	client := s3.NewFromConfig(*u.config)
	uploader := manager.NewUploader(client) // permite subir archivos concurrentemente
	_, e := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(pathWithName),
		Body:   r,
	})
	if e != nil {
		return nil, e
	}
	return &apstoreabstract.FileData{
		Path:       pathWithName,
		UploadedAt: time.Now().UTC().Unix(),
	}, nil
}

// Remove elimina un archivo en el bucket
func (u *AwsS3) Remove(pathFile string) error {
	client := s3.NewFromConfig(*u.config)
	_, e := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(u.bucket),
		Key:    aws.String(pathFile),
	})
	return e
}

// // Download descarga un archivo en el bucket de s3.
// func (u *AwsS3) Download(pathFile string, w io.WriterAt) error {

// 	// sess, err := session.NewSessionWithOptions(session.Options{
// 	// 	Config: aws.Config{
// 	// 		Region: aws.String(u.region),
// 	// 		// Credentials: credentials.NewStaticCredentials(u.accessKey, u.secretAccessKey, u.token),
// 	// 	},
// 	// })
// 	// if err != nil {
// 	// 	return errors.New("Problema creando session. Error: " + err.Error())
// 	// }
// 	// Creando el uploader con la session
// 	downloader := s3manager.NewDownloader(u.sess)

// 	// Subiendo el archivo al bucket de S3
// 	_, err2 := downloader.Download(w, &s3.GetObjectInput{
// 		Bucket: aws.String(u.bucket),
// 		Key:    aws.String(pathFile),
// 	})
// 	if err2 != nil {
// 		return fmt.Errorf("Fallo la descarga del archivo, %v", err2)
// 	}
// 	return nil
// }

// // Remove elimina un archivo en el bucket de s3.
// //
// // Retorna un error si ocurriera.
// // Retorna nil si se elimino el archivo o si el archivo no se encontro.
// func (u *AwsS3) Remove(pathFile string) error {

// 	if len(u.bucket) == 0 {
// 		return errors.New("Proporcione el nombre del bucket")
// 	}
// 	// sess, err := session.NewSessionWithOptions(session.Options{
// 	// 	Config: aws.Config{
// 	// 		Region: aws.String(u.region),
// 	// 		// Credentials: credentials.NewStaticCredentials(u.accessKey, u.secretAccessKey, u.token),
// 	// 	},
// 	// })
// 	// if err != nil {
// 	// 	return errors.New("Problema creando session. Error: " + err.Error())
// 	// }

// 	svc := s3.New(u.sess)

// 	// ****
// 	// **** Peticion a aws-s3 q comprobaria si existe el archivo a eliminar
// 	// // comprobando existencia del archivo a eliminar en el bucket de s3
// 	// input := &s3.HeadObjectInput{
// 	// 	Bucket: aws.String(u.bucket),
// 	// 	Key:    aws.String(pathFile),
// 	// }
// 	// _, err = svc.HeadObject(input)
// 	// // si ocurre error lo retorna.
// 	// // Incluidendo el error si no se encuentra el archivo a eliminar proposito de esta peticion
// 	// if err != nil {
// 	// 	if aerr, ok := err.(awserr.Error); ok {
// 	// 		return errors.New(aerr.Error())
// 	// 	}
// 	// 	return err
// 	// }
// 	// ****
// 	// **** FIN.

// 	input := &s3.DeleteObjectInput{
// 		Bucket: aws.String(u.bucket),
// 		Key:    aws.String(pathFile),
// 	}
// 	_, err := svc.DeleteObject(input)
// 	if err != nil {
// 		if aerr, ok := err.(awserr.Error); ok {
// 			return errors.New(aerr.Error())
// 		}
// 		return err
// 	}

// 	// *****
// 	return nil
// }
