package role

import (
	"ginCms/models"
)

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
