package collection

import (
	"context"

	"github.com/tinkler/moonmist/pkg/db"
	"github.com/tinkler/moonmist/pkg/mst"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c *Collection[T]) Fetch(ctx context.Context) error {
	if c.Page < 1 {
		c.Page = 1
	}
	se := db.Get(ctx).Session(&gorm.Session{Context: ctx})
	if len(c.Where) > 0 {
		if expr := createParser(se, nil).parse(c.Where); expr != nil {
			se.Statement.AddClause(clause.Where{Exprs: []clause.Expression{expr}})
		}
	}
	if !c.Opt().NoCount {
		copiedClause := se.Statement.Clauses
		if err := se.Model(&c.List).Count(&c.ItemCount).Error; err != nil {
			return mst.InternalServerError(err)
		}
		for k, c := range copiedClause {
			se.Statement.Clauses[k] = c
		}
	}

	se = c.Order(se)
	if c.Limit() > 0 {
		se = se.Limit(int(c.Limit())).Offset(int(c.Offset()))
	}

	if err := se.Find(&c.List).Error; err != nil {
		return mst.InternalServerError(err)
	}

	return nil
}
