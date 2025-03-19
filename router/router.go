package router

import (
	"tryweb/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.GET("/get", controller.GetUser)
	user.GET("/all", controller.GetAllUser)
	user.GET("/add", controller.AddUser)
}
