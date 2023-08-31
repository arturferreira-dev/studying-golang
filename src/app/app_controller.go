package app

import (
	"net/http"
	"time"

	"github.com/arturferreira-dev/studying-golang/src/shared/generics"
	"github.com/gin-gonic/gin"
)

type PostBody struct {
	Name *string   `json:"name"`
	Time time.Time `json:"time"`
}
type UpdateBody struct {
	Name string `json:"name"`
}

type AppController struct {
}

func (c AppController) Create(ctx *gin.Context) {
	var body PostBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, gin.H{"create": "ok", "body": body})
}

func (c AppController) List(ctx *gin.Context) {
	a := generics.Paginate{1, 2, 1, []interface{}{"kkkk", "kkk2"}}
	ctx.JSON(200, a.GetResult())
}

func (c AppController) Detail(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"Detail": "ok", "id": ctx.Params.ByName("id")})
}

func (c AppController) Update(ctx *gin.Context) {
	var body UpdateBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, gin.H{"Update": "ok", "id": ctx.Params.ByName("id"), "updateBody": body})
}

func (c AppController) Delete(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"Delete": "ok", "id": ctx.Params.ByName("id")})
}
