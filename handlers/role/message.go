package role

import (
	"ginCms/models"
)

// RestoreReq 还原的请求消息
type RestoreReq struct {
	IDList []uint `form:"id" binding:"gt=0"`
}

// DeleteReq 删除的请求消息
type DeleteReq struct {
	IDList []uint `form:"id" binding:"gt=0"`
	Force  bool   `form:"force" binding:""`
}

// 添加请求消息
type AddReq struct {
	models.Role
	// 需要额外校验的字段
	Title string `json:"title" binding:"required"`
	Key   string `json:"key" binding:"required"`
}

// AddReq to Role
func (req AddReq) ToRole() *models.Role {
	row := req.Role
	row.Title = req.Title
	row.Key = req.Key
	return &row
}

// GetRowReq GetRow接口的请求消息类型
type GetRowReq struct {
	// required, 必须的
	// gt, 数值 > 0
	ID uint `form:"id" binding:"required,gt=0"`
}

// GetListReq GetList请求参数类型
type GetListReq struct {
	// 过滤
	models.RoleFilter
	// 排序
	models.Sorter
	// 翻页
	models.Pager
}

// Clean 查询列表参数清理
func (req *GetListReq) Clean() {
	req.RoleFilter.Clean()
	req.Sorter.Clean()
	req.Pager.Clean()
}
