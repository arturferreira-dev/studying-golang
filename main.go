package main

import (
	"github.com/arturferreira-dev/studying-golang/src/app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	app.Router(r)
	r.Run()
}
