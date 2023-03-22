package querying

import (
	"fmt"
	"strings"

	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// Filter is a slice of strings that represents a filter.
type Filter []string

// AddFilter returns a new Filter with the given field, operator and value.
func AddFilter(field string, operator string, value any) querying.Filterer {
	return Filter{}.Add(field, operator, value)
}

// Add implements Filterer.
func (f Filter) Add(field string, operator string, value any) querying.Filterer {
	f = append(f, fmt.Sprintf("%s=%s,%v", field, operator, value))
	return f
}

// Field implements Filterer.
func (f Filter) Field(idx int) string {
	if f[idx] == "" {
		return ""
	}

	parts := strings.Split(f[idx], "=")
	return parts[0]
}

// Operator implements Filterer.
func (f Filter) Operator(idx int) string {
	if f[idx] == "" {
		return ""
	}

	parts := strings.Split(f[idx], "=")
	subparts := strings.SplitN(parts[1], ",", 2)
	return subparts[0]
}

// Slice implements Filterer.
func (f Filter) Slice() []string {
	return f
}

// Value implements Filterer.
func (f Filter) Value(idx int) string {
	if f[idx] == "" {
		return ""
	}

	parts := strings.Split(f[idx], "=")
	subparts := strings.SplitN(parts[1], ",", 2)
	return subparts[1]
}

// FilterScope returns a GORM scope that applies filtering to the query.
func FilterScope(filter querying.Filterer) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter == nil {
			return db
		}

		for i := range filter.Slice() {
			fmt.Println(filter.Field(i), filter.Value(i), filter.Operator(i))
			if filter.Field(i) == "" || filter.Value(i) == "" || filter.Operator(i) == "" {
				continue
			}

			columName := schema.NamingStrategy{}.ColumnName("", filter.Field(i))

			value := filter.Value(i)
			if filter.Operator(i) == "like" {
				value = fmt.Sprintf("%%%s%%", value)
			}

			if expr, ok := filterOperators[filter.Operator(i)]; ok {
				db = db.Where(expr(columName, value))
			}
		}

		return db
	}
}

const (
	EqOperator   = "eq"
	NeqOperator  = "neq"
	GtOperator   = "gt"
	GteOperator  = "gte"
	LtOperator   = "lt"
	LteOperator  = "lte"
	LikeOperator = "like"
	InOperator   = "in"
)

// filterOperators is a map that contains all allowed operators for filter parameters.
var filterOperators = map[string]func(column, value string) clause.Expression{
	EqOperator: func(column, value string) clause.Expression {
		return clause.Eq{Column: clause.Column{Name: column}, Value: value}
	},

	NeqOperator: func(column, value string) clause.Expression {
		return clause.Neq{Column: clause.Column{Name: column}, Value: value}
	},

	GtOperator: func(column, value string) clause.Expression {
		return clause.Gt{Column: clause.Column{Name: column}, Value: value}
	},

	GteOperator: func(column, value string) clause.Expression {
		return clause.Gte{Column: clause.Column{Name: column}, Value: value}
	},

	LtOperator: func(column, value string) clause.Expression {
		return clause.Lt{Column: clause.Column{Name: column}, Value: value}
	},

	LteOperator: func(column, value string) clause.Expression {
		return clause.Lte{Column: clause.Column{Name: column}, Value: value}
	},

	LikeOperator: func(column, value string) clause.Expression {
		return clause.Like{Column: clause.Column{Name: column}, Value: value}
	},

	InOperator: func(column, value string) clause.Expression {
		// Slice will be contained in a string, like so: "[1 2 3]".
		// We need to remove the brackets and split the string by space.
		values := strings.Split(strings.Trim(value, "[]"), " ")

		var anyValues []any
		for _, v := range values {
			anyValues = append(anyValues, any(v))
		}

		return clause.IN{Column: clause.Column{Name: column}, Values: anyValues}
	},
}
