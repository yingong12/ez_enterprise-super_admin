package controller

import "github.com/gin-gonic/gin"

func Health(ctx *gin.Context) {
	ctx.Data(200, "text", []byte("ok"))
}
