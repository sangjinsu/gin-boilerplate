package user

import (
	"gin-boiler-plate/controller"
	"github.com/gin-gonic/gin"
)

func SetUserGroup(app *gin.Engine) {
	user := controller.NewUserController()

	userGroup := app.Group("/v1/user")
	userGroup.GET(":id", user.FindById)
	userGroup.GET("all", user.FindAll)
	userGroup.POST("", user.Create)
	userGroup.PATCH(":id", user.Update)
	userGroup.DELETE(":id", user.Delete)
}
