package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	aphv1resp "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/resp"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

// response ...
type response struct {
	env aputil.EnvType
	c   *gin.Context
}

// send ...
func (r *response) send(httpCode int, code aphv1resp.CodeType, data interface{}) {
	if httpCode < 400 {
		r.c.JSON(httpCode, data)
	} else {
		msg := aphv1resp.GetMsgError(code) // mensaje por defecto
		// solo si esta en modo desarrollo se envia el error con todos sus detalles
		if r.env == aputil.EnvDevelop {
			switch t := data.(type) {
			case string:
				msg = fmt.Sprintf("%s --- %s", msg, t)
			case error:
				msg = fmt.Sprintf("%s --- %s", msg, t.Error())
			default:
				msg = fmt.Sprintf("%s --- %+v", msg, data)
			}
		}
		r.c.JSON(httpCode, &aphv1resp.Error{
			Code: code,
			Msg:  msg,
		})
	}
}

// sendOK ...
func (r *response) sendOK(data interface{}) {
	r.send(http.StatusOK, aphv1resp.CodeOK, data)
}

// sendBadRequest ...
func (r *response) sendBadRequest(code aphv1resp.CodeType, e error) {
	r.send(http.StatusBadRequest, code, e.Error())
}

// sendNotFound ...
func (r *response) sendNotFound(code aphv1resp.CodeType, e error) {
	r.send(http.StatusNotFound, code, e.Error())
}

// sendInternalError ...
func (r *response) sendInternalError(code aphv1resp.CodeType, e error) {
	r.send(http.StatusInternalServerError, code, e.Error())
}
