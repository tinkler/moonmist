package collection

import (
	"strings"

	"gorm.io/gorm"
)

func (c Collection[T]) Order(db *gorm.DB) *gorm.DB {
	for _, sort := range c.Sort {
		if len(sort) == 2 {
			sortOrder := strings.ToUpper(sort[1])
			if sortOrder != "ASC" && sortOrder != "DESC" {
				continue
			}
			db = db.Order(sort[0] + " " + sortOrder)
		} else if len(sort) == 1 {
			db = db.Order(sort[0] + " ASC")
		}
	}
	return db
}
