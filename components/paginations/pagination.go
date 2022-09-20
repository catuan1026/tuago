package paginations

var (
	defaultMaxPageSize = 100
)

func SetDefaultMaxPageSize(maxPageSize int) {
	defaultMaxPageSize = maxPageSize
}

func GetDefaultMaxPageSize() int {
	return defaultMaxPageSize
}

type Pagination[T any] struct {
	PageSize     int `json:"page_size"`
	PageNo       int `json:"page_no"`
	SearchOption T   `json:"search_option"`
}

func (p Pagination[T]) Limit() int {
	if p.PageSize <= 0 || p.PageSize > defaultMaxPageSize {
		return 10
	}
	return p.PageSize
}

func (p Pagination[T]) Offset() int {
	if p.PageNo <= 0 {
		p.PageNo = 1
	}
	return (p.PageNo - 1) * p.Limit()
}
