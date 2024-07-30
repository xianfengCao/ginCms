package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type FieldMap = map[string]any

// 查询范围的常量
const (
	SCOPE_ALL = iota
	SCOPE_UNDELETED
	SCOPE_DELETED
)

// 通用的查询列表排序类型
type Sorter struct {
	// 排序字段
	SortField *string `form:"sortField" binding:"omitempty,gt=0"`
	// 排序方式 asc,desc
	// oneof，多个选项之一
	SortMethod *string `form:"sortMethod" binding:"omitempty,oneof=asc desc"`
}

// 通用的查询列表翻页类型
type Pager struct {
	// 页码索引
	PageNum *int `form:"pageNum" binding:"omitempty,gt=0"`
	// 每页记录数
	PageSize *int `form:"pageSize" binding:"omitempty,gt=0"`
}

const (
	PageNumDefault  = 1
	PageSizeDefault = 10
	PageSizeMax     = 100

	SortFieldDefault  = "id"
	SortMethodDefault = "DESC"
)

// Clean 整理Sorter
func (s *Sorter) Clean() {
	if s.SortField == nil {
		temp := SortFieldDefault
		s.SortField = &temp
	}
	if s.SortMethod == nil {
		temp := SortMethodDefault
		s.SortMethod = &temp
	}
}

// Clean 整理Pager
func (p *Pager) Clean() {
	if p.PageNum == nil || *p.PageNum == 0 {
		temp := PageNumDefault
		p.PageNum = &temp
	}
	if p.PageSize == nil {
		temp := PageSizeDefault
		p.PageSize = &temp
	}
	if *p.PageSize > PageSizeMax {
		temp := PageSizeMax
		p.PageSize = &temp
	}
}
