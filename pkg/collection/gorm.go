package collection

import (
	"strings"

	"github.com/tinkler/moonmist/pkg/jsonz/sjson"
	"github.com/tinkler/moonmist/pkg/mlog"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type whereParser struct {
	db            *gorm.DB
	snakedNameMap map[string]string
}

func createParser(db *gorm.DB, snakedNameMap map[string]string) *whereParser {
	return &whereParser{
		db, snakedNameMap,
	}
}

func (p whereParser) parse(whereClause map[string]interface{}) clause.Expression {
	if len(whereClause) == 0 {
		return nil
	}

	var exprs []clause.Expression

	for k, v := range whereClause {
		upperKey := strings.ToUpper(k)
		switch upperKey {
		case "AND", "OR":
			clauseArr, ok := v.([]interface{})
			if !ok {
				mlog.Warn("And,Or require array")
				continue
			}
			var subExprs []clause.Expression
			for _, e := range clauseArr {
				subWhereClause, ok := e.(map[string]interface{})
				if !ok {
					mlog.Warn(upperKey + "'s list is not valid where clause")
					continue
				}
				if expr := p.parse(subWhereClause); expr != nil {
					subExprs = append(subExprs, expr)
				}
			}
			if len(subExprs) > 0 {
				if upperKey == "OR" {
					exprs = append(exprs, clause.Or(subExprs...))
				} else {
					exprs = append(exprs, clause.And(subExprs...))
				}
			}
		default:
			varName := strings.TrimSpace(k)
			if strings.Contains(varName, " ") {
				mlog.Warn("illegal where" + varName)
				continue
			}
			setValue, ok := v.(map[string]interface{})
			if !ok {
				continue
			}
			varName = sjson.ToSnakedName(varName)
			if p.snakedNameMap != nil {
				if name := p.snakedNameMap[varName]; name != "" {
					varName = name
				}
			}
			exprs = append(exprs, p.parseCompute(varName, setValue)...)
		}
	}

	if len(exprs) > 0 {
		return clause.And(exprs...)
	}
	return nil
}

func (p whereParser) parseCompute(varName string, setValue map[string]interface{}) (exprs []clause.Expression) {
	for k, v := range setValue {
		upperKey := strings.ToUpper(k)
		switch upperKey {
		case "EQ", "=":
			if !isNotArrayOrMap(v) {
				mlog.Warn("eq require value")
				continue
			}
			exprs = append(exprs, p.db.Statement.BuildCondition(varName+" = ?", v)...)
		case "NE", "<>":
			if !isNotArrayOrMap(v) {
				mlog.Warn("ne require value")
				continue
			}
			exprs = append(exprs, p.db.Statement.BuildCondition(varName+" <> ?", v)...)
		case ">", ">=", "<", "<=":
			if !isNotArrayOrMap(v) {
				mlog.Warn("ne require value")
				continue
			}
			exprs = append(exprs, p.db.Statement.BuildCondition(varName+" "+upperKey+" ?", v)...)
		case "IN":
			_, isArray := v.([]interface{})
			if !isArray {
				mlog.Warn("in require array")
				continue
			}
			exprs = append(exprs, p.db.Statement.BuildCondition(varName+" IN ?", v)...)
		case "IS":
			str, isString := v.(string)
			if !isString {
				mlog.Warn("is require null or not null")
				continue
			}
			switch strings.ToUpper(str) {
			case "NULL":
				exprs = append(exprs, p.db.Statement.BuildCondition(varName+" IS NULL")...)
			case "NOT NULL":
				exprs = append(exprs, p.db.Statement.BuildCondition(varName+" IS NOT NULL")...)
			}
		case "LIKE":
			str, isString := v.(string)
			if !isString {
				mlog.Warn("is require null or not null")
				continue
			}
			str = strings.TrimSpace(str)
			if strings.Contains(str, "%") {
				exprs = append(exprs, p.db.Statement.BuildCondition(varName+" LIKE ?", str)...)
			} else {
				exprs = append(exprs, p.db.Statement.BuildCondition(varName+" LIKE ?", "%"+str+"%")...)
			}
		case "BETWEEN":
			arr, isArray := v.([]interface{})
			if !isArray {
				mlog.Warn("between require array")
				continue
			}
			if len(arr) != 2 {
				mlog.Warn("between require a array with at less two value")
				continue
			}
			exprs = append(exprs, p.db.Statement.BuildCondition(varName+" BETWEEN ? AND ?", arr[0], arr[1])...)
		}
	}
	return
}

func isNotArrayOrMap(v interface{}) bool {
	if _, ok := v.([]interface{}); ok {
		return false
	}
	if _, ok := v.(map[string]interface{}); ok {
		return false
	}
	return true
}
