package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"super_admin/library/env"
	"super_admin/providers"

	"github.com/gin-gonic/gin"
)

func EnterpriseUpdate(ctx *gin.Context) (res RawResponse, err error) {
	appID := ctx.Query("app_id")
	if appID == "" {
		res = []byte("参数校验失败,缺少app_id")
		return
	}
	clientBody := &bytes.Reader{}
	if err = func() error {
		body := ctx.Request.Body
		bodyMap := map[string]interface{}{}
		err := json.NewDecoder(body).Decode(&bodyMap)
		if err != nil {
			return fmt.Errorf("decode failed")
		}
		//添加uid
		j, _ := json.Marshal(bodyMap)
		clientBody = bytes.NewReader(j)
		return nil
	}(); err != nil {
		return
	}
	//发送请求
	client := providers.HttpClientCompanyService
	//从url里拿appID
	URL := fmt.Sprintf("%s/enterprise/%s", client.BaseURL, appID)
	request, err := http.NewRequest("PUT", URL, clientBody)
	if err != nil {
		return
	}
	rsp, err := client.Do(request)
	if err != nil {
		return
	}
	//解析rsp
	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	//
	res = data
	return
}
func FowardGroupRequest(ctx *gin.Context) {
	Proxy(ctx, env.GetStringVal("LB_COMPANY_SERVICE"))
	return
}

func ForwardEnterpriseRequest(ctx *gin.Context) {
	//企业更新。 转发
	if ctx.Request.URL.Path == "/enterprise/update" || ctx.Request.URL.Path == "enterprise/update" {
		STDWrapperRaw(EnterpriseUpdate)(ctx)
		return
	}
	Proxy(ctx, env.GetStringVal("LB_COMPANY_SERVICE"))
	return
}
