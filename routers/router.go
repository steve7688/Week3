package routers

import (
	"Week3/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouters() *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/v1")

	//auth
	v1.POST("/login", controller.Login)

	return r
}
