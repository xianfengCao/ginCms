package role

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	g := r.Group("role")
	// 查询一条
	g.GET("", GetRow) // GET /role?id=21
	// 查询多条
	g.GET("list", GetList) // GET /role/list?
	// 添加
	g.POST("", Add) // POST /role
	// 删除
	g.DELETE("", Delete) // DELETE /role?id=22&id=33&id=44
	// 查询回收站
	g.GET("recycle", Recycle) // GET /role/recycle?
	// 回收站还原
	g.PUT("restore", Restore) // PUT /role?id=22&id=33&id=44
	// 更新单条
	g.PUT(":id", Edit) // PUT /role/33
	// 更新多条的enabled
	g.PUT("/enabled", EditEnabled) // PUT /role?id=11&id=22
}
