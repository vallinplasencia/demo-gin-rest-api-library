package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	aphv1req "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/req"
	aphv1resp "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/resp"
)

// BookHandler incoming request for books
type BookHandler struct {
	*Base
}

// PostBook add a new book
func (h *BookHandler) PostBook(c *gin.Context) {
	resp := response{c: c, env: h.env}
	var e error
	d := aphv1req.CreateBook{}

	if e = c.ShouldBindWith(&d, binding.JSON); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e)
		return
	}
	item := toModelBook(&d)
	// validate category id
	if _, e := h.DB.Categories().Find(item.CategoryID); e != nil {
		if e == apdbabstract.ErrorNoItems {
			resp.sendNotFound(aphv1resp.CodeNotFoundCategory, errors.New("category id not found"))
		} else {
			resp.sendInternalError(aphv1resp.CodeInternalError, e)
		}
		return
	}
	id, e := h.DB.Books().Add(item)

	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e)
		return
	}
	resp.sendOK(&aphv1resp.ResponseID{ID: id})
}
