package role

import (
	"ginCms/models"
	"ginCms/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func EditEnabled(ctx *gin.Context) {
	// 1. 解析Query请求数据
	query := EditEnabledQueryReq{}
	if err := ctx.ShouldBindQuery(&query); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}
	//log.Println(query)

	// 2. 解析Body请求数据
	body := EditEnabledBodyReq{}
	if err := ctx.ShouldBind(&body); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}
	//log.Println(body)

	// 3. 更新数据
	rowsNum, err := models.RoleUpdateEnabled(query.IDList, body.Enabled)
	if err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "数据更新错误",
		})
		return
	}

	// 4. 响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsNum,
	})
}

func Edit(ctx *gin.Context) {
	// 1. 解析URI请求数据
	uri := EditUriReq{}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}
	//log.Println(uri)

	// 2. 解析Body请求数据
	body := EditBodyReq{}
	if err := ctx.ShouldBind(&body); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}
	//log.Println(body)

	// 3. req to map
	fieldMap := body.ToFieldMap()
	//log.Println(fieldMap)
	if err := models.RoleUpdates(fieldMap, uri.ID); err != nil {
		// 记录日志
		log.Println(err)
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "数据更新错误",
		})
		return
	}

	// 3. 响应
	// 往往需要重新查询一边，获取最新的role信息
	row, err := models.RoleFetch(uri.ID, false)
	if err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "查询错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 100,
		"data": row,
	})
}

func Restore(ctx *gin.Context) {
	// 1. 解析请求数据
	req := RestoreReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}

	// 2. 还原数据
	rowsNum, err := models.RoleRestore(req.IDList)
	if err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "数据还原错误",
		})
		return
	}

	// 3. 响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsNum,
	})
}

func Delete(ctx *gin.Context) {
	// 1. 解析请求数据
	req := DeleteReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}
	//log.Println(req)

	// 2. 删除数据
	rowsNum, err := models.RoleDelete(req.IDList, req.Force)
	if err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "数据删除错误",
		})
		return
	}

	// 3. 响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowsNum,
	})
}

func Recycle(ctx *gin.Context) {
	list(ctx, models.SCOPE_DELETED, false)
}
func GetList(ctx *gin.Context) {
	list(ctx, models.SCOPE_UNDELETED, true)
}
func list(ctx *gin.Context, scope uint8, assoc bool) {
	// 1. 解析请求消息
	req := GetListReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}

	// 2. 整理请求参数
	req.Clean()

	// 3. 基于model查询
	rows, err := models.RoleFetchList(req.RoleFilter, req.Sorter, req.Pager, scope, assoc)
	if err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "查询错误",
		})
		return
	}

	// 4. 响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rows,
	})

}

func Add(ctx *gin.Context) {
	// 1. 解析请求数据
	req := AddReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}

	// 2. 利用模型完成插入
	role := req.ToRole()
	if err := models.RoleInsert(role); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "数据插入错误",
		})
		return
	}

	// 3. 响应
	// 往往需要重新查询一边，获取最新的role信息
	row, err := models.RoleFetch(role.ID, false)
	if err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "查询错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 100,
		"data": row,
	})
}

func GetRow(ctx *gin.Context) {
	// 1. 解析请求数据（消息）
	req := GetRowReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}

	// 2. 利用模型完成查询
	row, err := models.RoleFetch(req.ID, false)
	if err != nil {
		// 记录日志
		utils.Logger().Error(err.Error())
		// 直接响应
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "查询错误",
		})
		return
	}

	// 3. 响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": row,
	})
}
