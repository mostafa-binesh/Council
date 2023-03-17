package filters

import (
	D "docker/database"
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int    `json:"limit" query:"limit"` // per page
	Page       int    `json:"page" query:"page"`   // current page
	Sort       string `json:"sort" query:"sort"`   // !
	TotalRows  int64  `json:"totalRows"`           // total records
	TotalPages int    `json:"totalPages"`          // total pages
	// Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
func Paginate(value interface{}, pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	D.DB().Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
