package routes

import (
	"crud/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/id/:userId", controller.FindMembroById)
	r.GET("/user/email/:userEmail", controller.FindMembroByEmail)
	r.POST("/user", controller.CreateMembroController)
	r.PUT("/user", controller.AtualizaMembro)
	r.DELETE("/user/:userId", controller.DeleteMembro)
}
