package mewsql

import (
	"fmt"
	"strings"
)

const (
	selectOptionWhere = iota
	selectOptionOrderBy
	selectOptionLimit
	selectOptionOffset

	numSelectOptionKind
)

type SelectOption struct {
	kind                      int
	statement                 string
	statementWithPlaceholders []string
	placeholderSeparator      string
	bindVars                  []interface{}
}

func Select(
	columns string,
	table string,
	options ...SelectOption,
) (sql string, bindVars []interface{}) {
	optionBuffer := make([]string, numSelectOptionKind)

	numBindVars := 0
	for _, option := range options {
		if len(option.statementWithPlaceholders) == 0 {
			optionBuffer[option.kind] = option.statement
			continue
		}

		var buffer []string
		for idx, statement := range option.statementWithPlaceholders {
			numBindVars++
			buffer = append(
				buffer,
				strings.ReplaceAll(
					statement,
					"?",
					fmt.Sprintf("$%d", numBindVars),
				),
			)
			bindVars = append(bindVars, option.bindVars[idx])
		}

		optionBuffer[option.kind] = fmt.Sprintf(
			"%s %s",
			option.statement,
			strings.Join(
				buffer,
				option.placeholderSeparator,
			),
		)
	}

	sql = fmt.Sprintf(
		"SELECT %s FROM %s %s",
		columns,
		table,
		strings.Join(optionBuffer, " "),
	)
	return
}

func WithOrderBy(expression string, ascDesc string) SelectOption {
	ascDesc = strings.ToUpper(ascDesc)
	if ascDesc != "ASC" && ascDesc != "DESC" {
		ascDesc = ""
	}
	return SelectOption{
		kind:      selectOptionOrderBy,
		statement: fmt.Sprintf("ORDER BY %s %s", expression, ascDesc),
	}
}

func WithLimit(count int) SelectOption {
	return SelectOption{
		kind:      selectOptionLimit,
		statement: fmt.Sprintf("LIMIT %d", count),
	}
}

func WithOffset(count int) SelectOption {
	return SelectOption{
		kind:      selectOptionOffset,
		statement: fmt.Sprintf("OFFSET %d", count),
	}
}

type Condition struct {
	sqlQuery string
	bindVar  interface{}
}

func WithWhere(
	conditions ...Condition,
) SelectOption {
	res := SelectOption{
		kind:                 selectOptionWhere,
		placeholderSeparator: " AND ",
	}

	if len(conditions) > 0 {
		res.statement = "WHERE"
	}

	for _, condition := range conditions {
		res.statementWithPlaceholders = append(
			res.statementWithPlaceholders,
			condition.sqlQuery,
		)
		res.bindVars = append(res.bindVars, condition.bindVar)
	}

	return res
}

func WithCondition(sqlQuery string, bindVar interface{}) Condition {
	return Condition{sqlQuery: sqlQuery, bindVar: bindVar}
}
