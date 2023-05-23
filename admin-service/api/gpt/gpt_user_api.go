package gpt

import (
	"admin-service/entity"
	"admin-service/global"
	"admin-service/model/gpt"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"net/url"
)

func (api *GptUserApi) Register() {
	api.RouterGroup.POST("SyncAccount", api.SyncAccount)
	api.RouterGroup.POST("addExt", api.addExt)
}

type ReqBody struct {
	DevId string `json:"devId"`
	Agent string `json:"agent"`
}

func (api *GptUserApi) SyncAccount(c *gin.Context) {
	reqBody := new(ReqBody)
	//if err := c.BindJSON(&reqBody); err != nil {
	//	return
	//}
	if error := c.ShouldBindBodyWith(&reqBody, binding.JSON); error == nil {
		fmt.Println(reqBody)
	}
	result := requestSyncAccount(*reqBody)
	c.JSON(http.StatusOK, result)
}

func (api *GptUserApi) addExt(c *gin.Context) {
	ret := new(entity.ApiRet)
	model := new(gpt.GptUser)
	if ret.Error = c.ShouldBindBodyWith(&model, binding.JSON); ret.Error == nil {
		global.GCS_DB.FirstOrCreate(model)
		if result := global.GCS_DB.Save(model); result.Error == nil {
			ret.Ok = true
		}
	}
	ret.Data = model
	c.JSON(http.StatusOK, ret)
}

func requestSyncAccount(reqBody ReqBody) (result map[string]interface{}) {
	apiURL := "http://c167chatgpt.f129rqisuw8sz.top:8092/user/setWebuser"
	params := url.Values{}
	params.Add("apikey", "c6e182a778c0499d975d")
	//params.Add("devId", "3af2ffd5153fd72d")
	params.Add("devId", reqBody.DevId)
	params.Add("username", "gpt")
	params.Add("nickname", "gpt")
	//params.Add("agent", "DL0011")
	params.Add("agent", reqBody.Agent)
	fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error:", err)
		result["ok"] = false
		return result
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&result); err != nil {
		fmt.Println("Error decoding JSON:", err)
		result["ok"] = false
		return result
	}
	//fmt.Println(result)
	//if result["resultCode"] == 1 {
	//	result["ok"] = true
	//} else {
	//	result["ok"] = false
	//}
	return result
}
