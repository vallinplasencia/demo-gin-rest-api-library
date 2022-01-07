package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	aphv1resp "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/models/resp"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

// response ...
type response struct {
	env aputil.EnvType
	c   *gin.Context
}

// send ...
func (r *response) send(httpCode int, code aphv1resp.CodeType, data interface{}, headers map[string]string, abort bool) {
	for k, v := range headers {
		r.c.Header(k, v)
	}
	if httpCode < 400 {
		if data != nil {
			r.c.JSON(httpCode, data)
		} else {
			r.c.JSON(httpCode, nil)
		}
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
				if data != nil {
					msg = fmt.Sprintf("%s --- %+v", msg, data)
				}
			}
		}
		r.c.JSON(httpCode, &aphv1resp.Error{
			Code: code,
			Msg:  msg,
		})
	}
	if abort {
		r.c.Abort()
	}
}

// sendOK ...
func (r *response) sendOK(data interface{}, headers map[string]string, abort bool) {
	r.send(http.StatusOK, aphv1resp.CodeOK, data, headers, abort)
}

// sendBadRequest ...
func (r *response) sendBadRequest(code aphv1resp.CodeType, e error, headers map[string]string, abort bool) {
	r.send(http.StatusBadRequest, code, e, headers, abort)
}

// sendNotFound ...
func (r *response) sendNotFound(code aphv1resp.CodeType, e error, headers map[string]string, abort bool) {
	r.send(http.StatusNotFound, code, e, headers, abort)
}

// sendInternalError ...
func (r *response) sendInternalError(code aphv1resp.CodeType, e error, headers map[string]string, abort bool) {
	r.send(http.StatusInternalServerError, code, e, headers, abort)
}
