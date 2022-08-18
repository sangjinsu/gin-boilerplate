package router

import (
	"gin-boiler-plate/router/user"
	"github.com/gin-gonic/gin"
)

func SetRoutes(app *gin.Engine) {
	user.SetUserGroup(app)
}
