package routes

import (
	controllerAdmin "crud/src/controller/admin-controller"
	controller "crud/src/controller/membro-controller"

	"github.com/gin-gonic/gin"
)

func InitRoutesMembros(r *gin.RouterGroup, userController controller.MembroControllerInterface) {
	r.GET("/user/id/:userId", userController.FindMembroById)
	r.GET("/user/email/:userEmail", userController.FindMembroByEmail)
	r.POST("/user", userController.CreateMembroController)
	r.PUT("/user", userController.AtualizaMembro)
	r.DELETE("/user/:userId", userController.DeleteMembro)
}

func InitRoutesAdmins(r *gin.RouterGroup, adminController controllerAdmin.AdminControllerInterface) {
	r.POST("/admin", adminController.CreateAdminController)
}
