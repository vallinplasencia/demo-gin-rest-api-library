package req

import (
	"strings"
)

// Paginator paginator for items list.(Query String)
type Paginator struct {
	// Page pagina en el paginado. Mayor o igual a 1.
	Page uint `form:"_page"`
	// Limit cantidad de item a devolver. Mayor o igual a 1
	Limit uint `form:"_limit"`
}

type sortItem struct {
	Field      string
	Descendent bool
}

// Sort sort lits if items
type Sort struct {

	// Pattern ...
	// formato de Pattern: fieldname1,-fieldname2,-fieldnamw3,...
	//  * fieldname: ascendent order by fieldname
	//  * -fieldname: descendent order by fieldname
	//
	// First item have priority
	Pattern string `form:"_sort"`
}

// Parse parse sort return slice with fieldname and order
//
// formato de sort:
//  * fieldname: ascendent order by fieldname
//  * -fieldname: descendent order by fieldname
//
// First item have priority
func (s *Sort) Parse() []*sortItem {
	strSort := strings.Trim(s.Pattern, " ")
	fieldsOp := strings.Split(strSort, ",")
	sorts := []*sortItem{}

	for _, v := range fieldsOp {
		if lv := len(v); lv > 0 {
			s := &sortItem{
				Field:      v,
				Descendent: false,
			}
			if v[0:1] == "-" { // operator - descendent order
				if lv > 1 { // -field
					s.Field = v[1:]
					s.Descendent = true
				} else {
					continue
				}
			}
			sorts = append(sorts, s)
		}
	}
	return sorts
}

// OperatorQueryType operator allow on query string for filter
type OperatorQueryType string

const (
	// OperatorQueryEq ...
	OperatorQueryEq OperatorQueryType = ""
	// OperatorQueryNotEqual ...
	OperatorQueryNotEqual OperatorQueryType = "not"
	// OperatorQueryLessThan ...
	OperatorQueryLessThan OperatorQueryType = "lt"
	// OperatorQueryGreatThan ...
	OperatorQueryGreatThan OperatorQueryType = "gt"
	// OperatorQueryLessThanEqual ...
	OperatorQueryLessThanEqual OperatorQueryType = "lte"
	// OperatorQueryGreatThanEqual ...
	OperatorQueryGreatThanEqual OperatorQueryType = "gte"
	// OperatorQueryRange ...
	OperatorQueryRange OperatorQueryType = "rng"
	// OperatorQueryContain ...
	OperatorQueryContain OperatorQueryType = "ctn"
	// OperatorQueryStartWith ...
	OperatorQueryStartWith OperatorQueryType = "startw"
	// OperatorQueryEndWith ...
	OperatorQueryEndWith OperatorQueryType = "endw"
)
