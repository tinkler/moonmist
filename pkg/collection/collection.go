package collection

type Collection[T any] struct {
	ItemCount int64
	PageSize  int64
	Page      int64
	List      []T
	Sort      [][]string             `json:",omitempty"`
	Summary   T                      `json:",omitempty"`
	Where     map[string]interface{} `json:",omitempty"`
	Options   *Options               `json:",omitempty"`
}

var emptyOption = Options{}

func (c *Collection[T]) Opt() *Options {
	if c.Options == nil {
		return &emptyOption
	}
	return c.Options
}

type Options struct {
	Refresh bool `json:",omitempty"`
	NoCount bool `json:",omitempty"`
}

func (c Collection[T]) Limit() int64 {
	if c.PageSize == 0 {
		c.PageSize = 10
	}
	return c.PageSize
}

func (c Collection[T]) Offset() int64 {
	return (c.Page - 1) * c.PageSize
}
