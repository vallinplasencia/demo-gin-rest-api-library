package mysql

import (
	// database/sql"

	"fmt"
	"strconv"

	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
)

// valueType posibles tipos del valor de un campo
type valueType int

const (
	// valueInt numero entero
	valueInt valueType = iota + 1
	// valueFloat numero real
	valueFloat
	// valueString cadena
	valueString
	// valueBoolen boolean
	valueBoolen
)

// toConditionMysql procesa los parametros de entrada y retorna la condicion de una clausula de mysql y el valor a asignar o un error.
//
// Ej retorna:
// 1- name=?
// 2- Juan
// 3- nil
func toConditionMysql(field string, values []string, valueType valueType, op apdbabstract.OperatorType) (string, interface{}, error) {
	value := values[0]
	var v interface{}
	var e error
	condition := ""
	// validating value type
	switch valueType {
	case valueInt:
		v, e = strconv.ParseInt(value, 10, 64)
	case valueFloat:
		v, e = strconv.ParseFloat(value, 64)
	case valueBoolen:
		var b bool
		b, e = strconv.ParseBool(value)
		if b {
			v = 1
		} else {
			v = 0
		}
	default:
		v = value
	}
	if e != nil {
		return "", nil, e
	}
	// parse operator to condition db and/or process value
	switch op {
	case apdbabstract.OperatorEqual:
		condition = fmt.Sprintf("%s=?", field)
	case apdbabstract.OperatorNotEqual:
		condition = fmt.Sprintf("%s<>?", field)
	case apdbabstract.OperatorLessThan:
		condition = fmt.Sprintf("%s<?", field)
	case apdbabstract.OperatorLessThanEqual:
		condition = fmt.Sprintf("%s<=?", field)
	case apdbabstract.OperatorGreatThan:
		condition = fmt.Sprintf("%s>?", field)
	case apdbabstract.OperatorGreatThanEqual:
		condition = fmt.Sprintf("%s>=?", field)
	case apdbabstract.OperatorRange:
		condition = fmt.Sprintf("%s>=? AND %s<=?", field, field)
	case apdbabstract.OperatorContain:
		if valueType != valueString {
			return "", nil, apdbabstract.ErrorOperatorInvalidForValueType
		}
		v = fmt.Sprintf("%s%s%s", "%", v, "%")
		condition = fmt.Sprintf("%s LIKE ?", field)
	case apdbabstract.OperatorStartWith:
		if valueType != valueString {
			return "", nil, apdbabstract.ErrorOperatorInvalidForValueType
		}
		v = fmt.Sprintf("%s%s", v, "%")
		condition = fmt.Sprintf("%s LIKE ?", field)
	case apdbabstract.OperatorEndWith:
		if valueType != valueString {
			return "", nil, apdbabstract.ErrorOperatorInvalidForValueType
		}
		v = fmt.Sprintf("%s%s", "%", v)
		condition = fmt.Sprintf("%s LIKE ?", field)
	default:
		return "", nil, apdbabstract.ErrorOperatorInvalid
	}
	return condition, v, nil
}
