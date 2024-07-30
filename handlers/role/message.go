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

// ToFieldMap 将请求转换为存在字段的map
func (req EditBodyReq) ToFieldMap() models.FieldMap {
	fieldMap := models.FieldMap{}

	// 利用reflect遍历结构体，进而判断是否为nil，进行map构建
	rType := reflect.TypeOf(req)
	rValue := reflect.ValueOf(req)
	for i, l := 0, rType.NumField(); i < l; i++ {
		// 获取每个结构的字段类型
		tf := rType.Field(i)
		fieldTag := tf.Tag.Get("field")
		// 未指定field tag的字段，不做自动化处理
		if fieldTag == "" {
			continue
		}
		vf := rValue.Field(i)
		// 字段的值不是nil
		if !vf.IsNil() {
			//// 特定的字段，也不做自动化处理
			//if fieldTag == "some_filed" {
			//	continue
			//}
			fieldMap[fieldTag] = vf.Elem().Interface()
		}
	}

	return fieldMap
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
