package collection

import (
	"context"

	"github.com/tinkler/moonmist/pkg/db"
	"github.com/tinkler/moonmist/pkg/mst"
)

func (c *Collection[T]) Fetch(ctx context.Context) error {
	if c.Page < 1 {
		c.Page = 1
	}

	if err := db.Get(ctx).Find(&c.List).Error; err != nil {
		return mst.InternalServerError(err)
	}
	return nil
}
