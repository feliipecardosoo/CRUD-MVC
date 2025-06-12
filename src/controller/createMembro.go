package controller

import (
	resterr "crud/src/configuration/rest-err"
	"crud/src/controller/model/request"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateMembro(c *gin.Context) {

	var membroRequest request.MembroRequest

	if err := c.ShouldBindJSON(&membroRequest); err != nil {
		restErr := resterr.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, erro=%s\n", err.Error()),
		)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(201, membroRequest)
}
