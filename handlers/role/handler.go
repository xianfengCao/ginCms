package role

import (
	"ginCms/models"
	"ginCms/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetList(ctx *gin.Context) {
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

	log.Println(*req.Keyword, *req.SortField, *req.SortMethod, *req.PageNum, *req.PageSize)

	// 3. 基于model查询
	rows, err := models.RoleFetchList(false, req.RoleFilter, req.Sorter, req.Pager)
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
	row, err := models.RoleFetchRow(false, "`id` = ?", req.ID)
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
