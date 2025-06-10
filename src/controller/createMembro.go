package controller

import (
	resterr "crud/src/configuration/rest-err"

	"github.com/gin-gonic/gin"
)

func CreateMembro(c *gin.Context) {
	err := resterr.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
