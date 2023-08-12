package db

import (
	"fmt"
	"strconv"
	"strings"
)

type Value interface {
	ToString() string
}
type IntValue int

func (i IntValue) ToString() string {
	return strconv.Itoa(int(i))
}

type StringValue string

func (s StringValue) ToString() string {
	return string(s)
}

type QueryBuilder struct {
	SelectFields []string
	TableName    string
	Conditions   []string
	queryType    string
}

func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder {
	if len(fields) == 0 {
		fields = []string{"*"}
	}
	qb.SelectFields = fields
	qb.queryType = "Select"
	return qb
}

func (qb *QueryBuilder) Table(name string) *QueryBuilder {
	qb.TableName = name
	return qb
}

func (qb *QueryBuilder) Where(field string, operator string, value Value) *QueryBuilder {
	if _, ok := value.(StringValue); ok {
		value = StringValue(fmt.Sprintf("'%s'", value))
	}
	qb.Conditions = append(qb.Conditions, fmt.Sprintf("`%s` %s %v", field, operator, value))
	return qb
}

func (qb *QueryBuilder) Build() string {
	var query string
	switch qb.queryType {
	case "Select":
		query = fmt.Sprintf("SELECT %s FROM %s", strings.Join(qb.SelectFields, ", "), qb.TableName)
		if len(qb.Conditions) > 0 {
			query += fmt.Sprintf(" WHERE %s", strings.Join(qb.Conditions, " AND "))
		}
	}
	return query
}
