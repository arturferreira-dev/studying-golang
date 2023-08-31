package generics

import (
	"github.com/arturferreira-dev/studying-golang/src/shared/interfaces"
	"github.com/gin-gonic/gin"
)

func GenericRouter(path string, r *gin.Engine, c interfaces.ControllerGeneric) {
	r.POST(path, c.Create)
	r.GET(path, c.List)
	r.GET(path+"/:id", c.Detail)
	r.PUT(path+"/:id", c.Update)
	r.DELETE(path+"/:id", c.Delete)
}
