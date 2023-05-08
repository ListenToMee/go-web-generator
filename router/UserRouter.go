package router

import (
	"github.com/gin-gonic/gin"
	"go-web-generator/controller"
)

// UserRouter 路由
func UserRouter(engin *gin.Engine) {

	// 定义路由组
	group := engin.Group("/user")

	// 绑定分页接口
	group.POST("/queryPage", controller.UserController{}.QueryPage)

	// 绑定新增接口
	group.PUT("/add", controller.UserController{}.Add)

	// 绑定更新接口
	group.PUT("/update", controller.UserController{}.Update)

	// 绑定查询接口
	group.GET("/queryDetails", controller.UserController{}.QueryDetails)

	// 绑定删除接口
	group.DELETE("/delete", controller.UserController{}.Delete)

	// 绑定批量删除接口
	group.DELETE("/deleteBatch", controller.UserController{}.DeleteBatch)
}
