package utils

import (
	"math"
)

type Paginate struct {
	TotalCount  int64       `json:"totalCount"`
	CurrentPage int         `json:"currentPage"`
	TotalPage   int         `json:"totalPage"`
	Limit       int         `json:"limit"`
	DataList    interface{} `json:"dataList"`
}

func NewPaginate(totalCount int64, page int, limit int, dataList interface{}) *Paginate {
	p := &Paginate{TotalCount: totalCount, CurrentPage: page, Limit: limit, DataList: dataList}
	p.page()
	return p
}

func (p *Paginate) page() *Paginate {
	p.TotalPage = int(math.Ceil(float64(p.TotalCount) / float64(p.Limit)))
	return p
}
