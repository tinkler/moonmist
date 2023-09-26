package collection

type Collection[T any] struct {
	ItemCount int64
	PageSize  int64
	Page      int64
	List      []T
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
}
