package v1

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	aphv1req "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/req"
	aphv1resp "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/resp"
	apmodels "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
	apv1models "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// BookHandler incoming request for books
type BookHandler struct {
	*base
}

// PostAddBook add a new book
func (h *BookHandler) PostAddBook(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionAddBook) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, ErrorUnauthorized, true)
		return
	}
	var e error
	d := aphv1req.CreateBook{}

	if e = c.ShouldBindWith(&d, binding.JSON); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e, true)
		return
	}
	item := h.toModelBookFromRequest(&d, u.UserID)
	// validate category id
	if _, e := h.db.Categories().Find(item.CategoryID); e != nil {
		if e == apdbabstract.ErrorNoItems {
			resp.sendNotFound(aphv1resp.CodeNotFoundCategory, errors.New("category id not found"), true)
		} else {
			resp.sendInternalError(aphv1resp.CodeInternalError, e, true)
		}
		return
	}
	id, e := h.db.Books().Add(item)

	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, true)
		return
	}
	resp.sendOK(&aphv1resp.ResponseID{ID: id}, false)
}

// GetRetrieveBook get book
func (h *BookHandler) GetRetrieveBook(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionRetrieveBook) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, ErrorUnauthorized, true)
		return
	}
	item, e := h.db.Books().FindBookSimple(c.Param("id"))
	if e != nil {
		if e == apdbabstract.ErrorNoItems {
			resp.sendNotFound(aphv1resp.CodeNotFoundBook, e, true)
			return
		}
		resp.sendInternalError(aphv1resp.CodeInternalError, e, true)
		return
	}
	resp.sendOK(h.toRespBookFromDB(item), false)
}

// === conv from request === //

func (h *BookHandler) toModelBookFromRequest(d *aphv1req.CreateBook, userID string) *apv1models.Book {
	return &apv1models.Book{
		Title:       d.Title,
		Original:    d.Original,
		Tags:        d.Tags,
		PublishedAt: d.PublishedAt,
		CreatedAt:   time.Now().UTC().Unix(),
		UpdatedAt:   0,
		CategoryID:  d.CategoryID,
		UserID:      userID,
	}
}

// === conv from db === //

func (h *BookHandler) toRespBookFromDB(d *apv1models.BookSimple) *aphv1resp.BookSimple {
	c := d.Category
	return &aphv1resp.BookSimple{
		ID:          d.ID,
		Title:       d.Title,
		Original:    d.Original,
		Tags:        d.Tags,
		PublishedAt: d.PublishedAt,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		UserID:      d.UserID,
		Category: &aphv1resp.CategorySimple{
			ID:          c.ID,
			Name:        c.Name,
			Description: c.Description,
		},
	}
}
