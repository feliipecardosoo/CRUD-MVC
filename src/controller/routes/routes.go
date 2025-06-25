package routes

import (
	"crud/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.MembroControllerInterface) {

	r.GET("/user/id/:userId", userController.FindMembroById)
	r.GET("/user/email/:userEmail", userController.FindMembroByEmail)
	r.POST("/user", userController.CreateMembroController)
	r.PUT("/user", userController.AtualizaMembro)
	r.DELETE("/user/:userId", userController.DeleteMembro)
}
