package role

import (
	"ginCms/models"
	"reflect"
)

// EditEnabledQueryReq 更新enabled的请求消息
// 将全部的Enabled字段，设置为相同的值
type EditEnabledQueryReq struct {
	IDList []uint `form:"id" binding:"gt=0"`
}

type EditEnabledBodyReq struct {
	Enabled bool `json:"enabled"`
}

// 提供id和enabled的映射，将某个id的enabled设置为某个值
//type EditEnabledBodyReq struct {
//	Enabled map[uint]bool `json:"enabled"`
//}

// EditUriReq URI上的id参数
type EditUriReq struct {
	ID uint `uri:"id" binding:"required,gt=0"` // 可以考虑加一个id存在的校验
}

// EditBodyReq 更新主体参数
type EditBodyReq struct {
	Title   *string `json:"title" field:"title"`
	Key     *string `json:"key" field:"key"`
	Enabled *bool   `json:"enabled" field:"enabled"`
	Weight  *int    `json:"weight" field:"weight"`
	Comment *string `json:"comment" field:"comment"`
}

func (req EditBodyReq) ToFieldMap() models.FieldMap {
	// 1. 初始化map
	m := models.FieldMap{}

	// 2. 利用反射来遍历req结构的全部字段
	reqType := reflect.TypeOf(req)
	reqValue := reflect.ValueOf(req)
	// 通过字段数量，进行遍历
	for i, nums := 0, reqType.NumField(); i < nums; i++ {
		// 获取 field tag
		fieldTag := reqType.Field(i).Tag.Get("field")
		// 存在 field tag才自动处理
		if fieldTag == "" {
			continue
		}
		// 判断字段是否为nil，这个值的判断
		if !reqValue.Field(i).IsNil() {
			if fieldTag == "some_field" {
				// 考虑特殊字段的处理情况
			} else {
				// 放入map
				m[fieldTag] = reqValue.Field(i).Elem().Interface()
			}
		}
	}

	return m
}

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
