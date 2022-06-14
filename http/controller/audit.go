package controller

import (
	"super_admin/library/env"

	"github.com/gin-gonic/gin"
)

func Search(ctx *gin.Context) (res STDResponse, err error) {
	return
}
func Create(ctx *gin.Context) (res STDResponse, err error) {
	return
}

func FowardAuditRequest(ctx *gin.Context) {
	Proxy(ctx, env.GetStringVal("LB_COMPANY_SERVICE"))
	return
}
