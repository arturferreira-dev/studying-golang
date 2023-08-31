package app

import (
	"github.com/arturferreira-dev/studying-golang/src/shared/generics"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	generics.GenericRouter("/home", r, AppController{})
}
