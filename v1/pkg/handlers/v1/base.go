package v1

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

	apauthtokenabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth/access-token/abstract"
	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	apstoreabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/store/abstract"
	apv1models "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

var ErrorRespIsEmpty error = errors.New("response is nil")

const (
	concFindAccountByEmail    aputil.ConcName = "find-account-by-email"
	concFindAccountByUsername aputil.ConcName = "find-account-by-username"
)

const (
	// avatarDirectoryIn directorio interno donde se suben los avatars
	avatarDirectoryIn string = "xxxvvv/media/avatars"
)

// Base ...
type base struct {
	env aputil.EnvType

	// === services  === //

	// db service
	db apdbabstract.DB
	// upload service: AWS-S3 or files system
	storeFiles apstoreabstract.Store
	// access-token(jwt) and refresh-token for auth
	token apauthtokenabstract.Token
}

// === funciones comunes q pueden ser usadas en cualesquiera de los handlers === //

// findAccountByEmail recupera una cuenta por su email
func (b *base) findAccountByEmail(email string, ch chan *aputil.ConcurrencyData, out *apv1models.Account) {
	d, e := b.db.Accounts().FindByEmail(email)
	cd := &aputil.ConcurrencyData{
		Err:  nil,
		Name: concFindAccountByEmail,
	}
	switch {
	case e != nil:
		cd.Err = e
	case d == nil:
		cd.Err = ErrorRespIsEmpty
	default:
		*out = *d
	}
	ch <- cd
}

// findAccountByUsername recupera una cuenta por su username
func (b *base) findAccountByUsername(username string, ch chan *aputil.ConcurrencyData, out *apv1models.Account) {
	d, e := b.db.Accounts().FindByUsername(username)
	cd := &aputil.ConcurrencyData{
		Err:  nil,
		Name: concFindAccountByUsername,
	}
	switch {
	case e != nil:
		cd.Err = e
	case d == nil:
		cd.Err = ErrorRespIsEmpty
	default:
		*out = *d
	}
	ch <- cd
}

// saveUploadFile salva el archivo q se intenta subir
func (b *base) saveUploadFile(directoryIn string, file *multipart.FileHeader) (string, error) {
	directoryIn = strings.TrimSuffix(directoryIn, string(os.PathSeparator))
	yyyy, mm, dd := time.Now().Date()
	filename := file.Filename
	ext := filename[strings.LastIndex(filename, "."):]
	rand := aputil.RandStringn(17)
	fullpath := fmt.Sprintf("%s/%d/%s/%d/%s%s", directoryIn, yyyy, strings.ToLower(mm.String()), dd, rand, ext)
	r, e := file.Open()
	if e != nil {
		return "", e
	}
	fd, e := b.storeFiles.Save(fullpath, r)
	if e != nil {
		return "", e
	}
	return fd.Path, nil
}
