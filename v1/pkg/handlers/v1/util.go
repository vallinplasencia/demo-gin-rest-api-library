package v1

import (
	"errors"

	ua "github.com/mileusna/useragent"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	aphv1req "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/req"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

// === errors === //

var (
	// errorRespIsEmpty external service response is empty
	errorRespIsEmpty error = errors.New("response is nil")
	// errorUnauthorized user not have permisssion to a resource
	errorUnauthorized error = errors.New("user does not have permission")
	// errorFieldNotSort not sort by field
	errorFieldNotSort error = errors.New("field no sort")
	// errorExpiredRefreshToken expired refresh token
	errorExpiredRefreshToken error = errors.New("refresh token expired")
	// errorInvalidSessionDataOfRefreshToken session data asociated with the refresh token is invalid. Ex: diferente platform, ...
	errorInvalidSessionDataOfRefreshToken error = errors.New("session data of refresh token is/are invalid")
)

// === concurrency === //
const (
	concFindAccountByEmail    aputil.ConcName = "find-account-by-email"
	concFindAccountByUsername aputil.ConcName = "find-account-by-username"
)

// === headers === //

const (
	// headerXTotalItems header para el total de items en un listado
	headerXTotalItems string = "X-Total-Items"
)

// === others === //

const (
	// avatarDirectoryIn directorio interno donde se suben los avatars
	avatarDirectoryIn string = "xxxvvv/media/avatars"
)

// === db util === //

// toOperatorDBFromRequest
func toOperatorDBFromRequest(op string) apdbabstract.OperatorType {
	switch op {
	case string(aphv1req.OperatorQueryEq):
		return apdbabstract.OperatorEqual
	case string(aphv1req.OperatorQueryNotEqual):
		return apdbabstract.OperatorNotEqual
	case string(aphv1req.OperatorQueryLessThan):
		return apdbabstract.OperatorLessThan
	case string(aphv1req.OperatorQueryGreatThan):
		return apdbabstract.OperatorGreatThan
	case string(aphv1req.OperatorQueryLessThanEqual):
		return apdbabstract.OperatorLessThanEqual
	case string(aphv1req.OperatorQueryGreatThanEqual):
		return apdbabstract.OperatorGreatThanEqual
	case string(aphv1req.OperatorQueryRange):
		return apdbabstract.OperatorRange
	case string(aphv1req.OperatorQueryContain):
		return apdbabstract.OperatorContain
	case string(aphv1req.OperatorQueryStartWith):
		return apdbabstract.OperatorStartWith
	case string(aphv1req.OperatorQueryEndWith):
		return apdbabstract.OperatorEndWith
	}
	return apdbabstract.OperatorEqual
}

// === request util === //

// parseUserAgent parsea userAgentStr y retorna la informacion
func parseUserAgent(userAgentStr string) *userAgent {
	d := ua.Parse(userAgentStr)
	return &userAgent{
		Name:   d.Name,
		OS:     d.OS,
		Device: d.Device,
		String: d.String,
	}
}

type userAgent struct {
	Name   string
	OS     string
	Device string
	String string
}
