package filters

import "gorm.io/gorm"

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"size"`
}

func Paginate(p *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// if page was not set, set it to default (1)
		if p.Page <= 0 {
			p.Page = 1
		}
		// if pageSize was not set, set it to default (8)
		switch {
		case p.PageSize > 100:
			p.PageSize = 100
		case p.PageSize <= 0:
			p.PageSize = 8
		}
		offset := (p.Page - 1) * p.PageSize
		return db.Offset(offset).Limit(p.PageSize)
	}
}
