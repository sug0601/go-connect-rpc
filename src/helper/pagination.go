package helper

type Pagination struct {
	Page      int
	PageSize  int
	Offset    int
	Total     int
	TotalPage int
}

const (
	DefaultPage     = 1
	DefaultPageSize = 10
	MaxPageSize     = 100
)

func NewPagination(page, pageSize int) Pagination {
	if page <= 0 {
		page = DefaultPage
	}
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	offset := (page - 1) * pageSize

	return Pagination{
		Page:     page,
		PageSize: pageSize,
		Offset:   offset,
	}
}

func (p *Pagination) SetTotal(total int) {
	p.Total = total
	if p.PageSize > 0 {
		p.TotalPage = (total + p.PageSize - 1) / p.PageSize
	}
}
