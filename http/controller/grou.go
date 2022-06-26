package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"super_admin/library/env"
	"super_admin/providers"

	"github.com/gin-gonic/gin"
)

func GroupUpdate(ctx *gin.Context) (res RawResponse, err error) {

	clientBody := &bytes.Reader{}
	var appID string
	if err = func() error {
		body := ctx.Request.Body
		bodyMap := map[string]interface{}{}
		err := json.NewDecoder(body).Decode(&bodyMap)
		if err != nil {
			return fmt.Errorf("decode failed")
		}
		//校验app_id
		v, ok := bodyMap["app_id"]
		if !ok {
			return errors.New("参数校验失败,缺少app_id")
		}
		vs, okForce := v.(string)
		if !okForce {
			return errors.New("app_id需要为string")
		}
		if vs == "" {
			return errors.New("参数校验失败,缺少app_id")
		}
		appID = vs
		j, _ := json.Marshal(bodyMap)
		clientBody = bytes.NewReader(j)
		return nil
	}(); err != nil {
		//直接把错误抛给后端
		res = []byte(err.Error())
		return
	}
	//发送请求
	client := providers.HttpClientCompanyService
	//从url里拿appID
	URL := fmt.Sprintf("%s/group/%s", client.BaseURL, appID)
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
	if ctx.Request.URL.Path == "/group/update" || ctx.Request.URL.Path == "group/update" {
		STDWrapperRaw(GroupUpdate)(ctx)
		return
	}
	Proxy(ctx, env.GetStringVal("LB_COMPANY_SERVICE"))
	return
}
