package v1

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	aphv1req "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/req"
	aphv1resp "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/resp"
	apmodels "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
	apv1models "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
)

// bookFieldType ...
type bookFieldType string

// field names on query string of books
const (
	fieldBookQueryTitle    bookFieldType = "title"
	fieldBookQueryOriginal bookFieldType = "original"
	fieldBookQueryTag      bookFieldType = "tag"
	// fieldBookQueryTags         bookFieldType = "tags"
	fieldBookQueryPublishedAt  bookFieldType = "published-at"
	fieldBookQueryCategoryName bookFieldType = "category-name"
	fieldBookQueryCategoryID   bookFieldType = "category-id"
)

// BookHandler incoming request for books
type BookHandler struct {
	*base
}

// PostAddBook add a new book
func (h *BookHandler) PostAddBook(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionAddBook) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, errorUnauthorized, nil, true)
		return
	}
	var e error
	d := aphv1req.CreateBook{}

	if e = c.ShouldBindWith(&d, binding.JSON); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e, nil, true)
		return
	}
	item := h.toModelBookFromRequest(&d, u.UserID)
	// validate category id
	if _, e := h.db.Categories().Find(item.CategoryID); e != nil {
		if e == apdbabstract.ErrorNoItems {
			resp.sendNotFound(aphv1resp.CodeNotFoundCategory, e, nil, true)
			return
		}
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	id, e := h.db.Books().Add(item)

	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	resp.sendOK(&aphv1resp.ResponseID{ID: id}, nil, false)
}

// PutEditBook update book
func (h *BookHandler) PutEditBook(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionEditBook) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, errorUnauthorized, nil, true)
		return
	}
	var e error
	d := aphv1req.EditBook{}

	if e = c.ShouldBindWith(&d, binding.JSON); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e, nil, true)
		return
	}
	item, e := h.db.Books().Find(c.Param("id"))
	if e != nil {
		if e == apdbabstract.ErrorNoItems {
			resp.sendNotFound(aphv1resp.CodeNotFoundBook, e, nil, true)
			return
		}
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	if item.CategoryID != d.CategoryID {
		// validando el id de la categorya solo si es diferente a la q ya tenia el libro
		if _, e := h.db.Categories().Find(item.CategoryID); e != nil {
			if e == apdbabstract.ErrorNoItems {
				resp.sendNotFound(aphv1resp.CodeNotFoundCategory, e, nil, true)
				return
			}
			resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
			return
		}
	}
	// solo el usuario q agrego el libro puede editarlo
	if item.UserID != u.UserID {
		// se envia un NOTFOUND pq esta tratando de aceder un usuario q no es el propietario del libro
		resp.sendNotFound(aphv1resp.CodeNotFoundBook, nil, nil, true)
		return
	}
	e = h.db.Books().Edit(item)
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	resp.sendOK(struct{}{}, nil, false)
}

// GetRetrieveBook get book
func (h *BookHandler) GetRetrieveBook(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionRetrieveBook) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, errorUnauthorized, nil, true)
		return
	}
	item, e := h.db.Books().FindBookSimple(c.Param("id"))
	if e != nil {
		if e == apdbabstract.ErrorNoItems {
			resp.sendNotFound(aphv1resp.CodeNotFoundBook, e, nil, true)
			return
		}
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	// solo el usuario q agrego el libro puede obtener sus datos
	if item.UserID != u.UserID {
		// se envia un NOTFOUND pq esta tratando de aceder un usuario q no es el proietario del libro
		resp.sendNotFound(aphv1resp.CodeNotFoundBook, e, nil, true)
		return
	}
	resp.sendOK(h.toRespBookSimpleFromDB(item), nil, false)
}

// GetListBooks list books
func (h *BookHandler) GetListBooks(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionListBooks) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, errorUnauthorized, nil, true)
		return
	}
	var e error
	dp := aphv1req.Paginator{}
	ds := aphv1req.Sort{}
	// paginator
	if e = c.ShouldBindWith(&dp, binding.Query); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e, nil, true)
		return
	}
	// sort
	if e = c.ShouldBindWith(&ds, binding.Query); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e, nil, true)
		return
	}
	fieldsFilter := map[bookFieldType][]aphv1req.OperatorQueryType{
		fieldBookQueryTitle:       {aphv1req.OperatorQueryEq, aphv1req.OperatorQueryNotEqual, aphv1req.OperatorQueryNotEqual, aphv1req.OperatorQueryStartWith, aphv1req.OperatorQueryEndWith, aphv1req.OperatorQueryContain},
		fieldBookQueryOriginal:    {aphv1req.OperatorQueryEq},
		fieldBookQueryCategoryID:  {aphv1req.OperatorQueryEq},
		fieldBookQueryTag:         {aphv1req.OperatorQueryContain},
		fieldBookQueryPublishedAt: {aphv1req.OperatorQueryLessThanEqual, aphv1req.OperatorQueryGreatThanEqual},
	}
	f := h.filtersBooks(c.Request.URL.Query(), fieldsFilter)
	s, _ := h.sortBooks(&ds, []bookFieldType{fieldBookQueryTitle, fieldBookQueryOriginal, fieldBookQueryPublishedAt, fieldBookQueryCategoryName})

	// for k, v := range f.Items {
	// 	fmt.Printf("KK: %+v\n", k)
	// 	for _, v := range v {
	// 		fmt.Printf("---VV: %+v\n", v)
	// 	}
	// }
	item, e := h.db.Books().ListBooksSimple(dp.Page, dp.Limit, s, f)
	if e != nil {
		// if e == apdbabstract.ErrorNoItems {
		// 	resp.sendNotFound(aphv1resp.CodeNotFoundBook, e, true)
		// 	return
		// }
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	resp.sendOK(h.toRespBooksSimpleFromDB(item), nil, false)
}

// DeleteBook remove book
func (h *BookHandler) DeleteBook(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionDeleteBook) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, errorUnauthorized, nil, true)
		return
	}
	item, e := h.db.Books().FindBookSimple(c.Param("id"))
	if e != nil {
		if e == apdbabstract.ErrorNoItems {
			resp.sendNotFound(aphv1resp.CodeNotFoundBook, e, nil, true)
			return
		}
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	// check owner book
	if item.UserID != u.UserID {
		resp.sendNotFound(aphv1resp.CodeNotFoundBook, e, nil, true)
		return
	}
	e = h.db.Books().Remove(item.ID)
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	resp.sendOK(struct{}{}, nil, false)
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

func (h *BookHandler) toRespBookSimpleFromDB(d *apv1models.BookSimple) *aphv1resp.BookSimple {
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

func (h *BookHandler) toRespBooksSimpleFromDB(d []*apv1models.BookSimple) []*aphv1resp.BookSimple {
	items := []*aphv1resp.BookSimple{}
	for _, v := range d {
		items = append(items, h.toRespBookSimpleFromDB(v))
	}
	return items
}

// === sort on db === //

// sortBooks parse to db sort and validate field allow to sort
func (h *BookHandler) sortBooks(sort *aphv1req.Sort, allowFieldsSort []bookFieldType) (*apdbabstract.SortBooks, error) {
	field := fieldBookQueryTitle
	s := &apdbabstract.SortBooks{
		Field:      apdbabstract.BookFieldTitle,
		Descendent: false,
	}
	if sort := sort.Parse(); len(sort) > 0 {
		s.Descendent = sort[0].Descendent

		switch v := sort[0]; v.Field {
		case string(fieldBookQueryTitle):
			s.Field = apdbabstract.BookFieldTitle
			field = fieldBookQueryTitle
		case string(fieldBookQueryOriginal):
			s.Field = apdbabstract.BookFieldOriginal
			field = fieldBookQueryOriginal
		case string(fieldBookQueryPublishedAt):
			s.Field = apdbabstract.BookFieldPublishedAt
			field = fieldBookQueryPublishedAt
		case string(fieldBookQueryCategoryName):
			s.Field = apdbabstract.BookFieldCategoryName
			field = fieldBookQueryCategoryName
		}
	}
	// validating field allow for ordenation
	valid := false
	for _, v := range allowFieldsSort {
		if v == field {
			valid = true
			break
		}
	}
	if valid {
		return s, nil
	}
	return &apdbabstract.SortBooks{
		Field:      apdbabstract.BookFieldTitle,
		Descendent: false,
	}, errorFieldNotSort
}

// === filters on db === //

// filtersBooks parse to db filter and validate field and yours operators allow filter
//
// mapFilters containe field and values for fielter. Format: name[not]=pepe, age[lt]=30, gender=male,....
// allowFieldsFilters: field and yours operators valid for filters
func (h *BookHandler) filtersBooks(mapFilters map[string][]string, allowFieldsFilters map[bookFieldType][]aphv1req.OperatorQueryType) *apdbabstract.FilterBooks {
	items := map[apdbabstract.BookFieldType][]*apdbabstract.FilterBooksItem{}

	for k, v := range mapFilters {
		var field bookFieldType

		if len(v) == 0 { // si no tiene values
			continue
		}
		param := k
		opss := ""
		op := apdbabstract.OperatorEqual
		// find operator between square
		if posFirstSquare := strings.Index(k, "["); posFirstSquare != -1 {
			param = k[:posFirstSquare]
			opss = k[posFirstSquare+1 : len(k)-1]
			op = toOperatorDBFromRequest(opss)
		}
		f := &apdbabstract.FilterBooksItem{
			Operator: op,
			Value:    v[0],
			Values:   v,
		}
		// parse field for db field
		switch param {
		case string(fieldBookQueryTitle):
			f.Field = apdbabstract.BookFieldTitle
			field = fieldBookQueryTitle
		case string(fieldBookQueryOriginal):
			f.Field = apdbabstract.BookFieldOriginal
			field = fieldBookQueryOriginal
		case string(fieldBookQueryPublishedAt):
			f.Field = apdbabstract.BookFieldPublishedAt
			field = fieldBookQueryPublishedAt
		case string(fieldBookQueryCategoryID):
			f.Field = apdbabstract.BookFieldCategoryID
			field = fieldBookQueryCategoryName
		case string(fieldBookQueryTag):
			f.Field = apdbabstract.BookFieldTags
			field = fieldBookQueryTag
		default:
			// si el nombre del parametro NO esta permitido no se agrega el filtro
			continue
		}
		// validate field allow and yours operators
		validField := false
		validOperator := false
		if ops, ok := allowFieldsFilters[field]; ok { // check valid field
			validField = true
			for _, v := range ops {
				if opss == string(v) { // check valid operator for field
					validOperator = true
					break
				}
			}
		}
		// check valid field and yours valid operator
		if !validField || !validOperator {
			continue
		}
		if values, ok := items[f.Field]; ok {
			items[f.Field] = append(values, f)
		} else {
			items[f.Field] = []*apdbabstract.FilterBooksItem{f}
		}
	}
	return &apdbabstract.FilterBooks{
		Items: items,
	}
}
