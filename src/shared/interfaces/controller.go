package interfaces

import (
	"github.com/gin-gonic/gin"
)

type ControllerGeneric interface {
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
