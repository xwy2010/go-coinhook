package controller

import (
	"encoding/json"
	o "hrforceAdmin/utils/apioutput"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type result struct {
	Code string `json:"code"`
	Data string `json:"data"`
}

// UpdateToken 更新微信token接口
func UpdateToken(c echo.Context) (err error) {
	result := &result{}
	resp, err := http.Get("http://172.16.14.84:10086/v1/server/wxtoken.update")
	//resp, err :=   http.Get("https://stage.hrforce.cn/api/v1/server/wxtoken.update")
	if err != nil {
		return c.JSON(http.StatusOK, o.ResultFail("刷新失败"))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusOK, o.ResultFail("刷新失败"))
	}
	err = json.Unmarshal(body, result) //JSON TO Struct
	if err != nil {
		return c.JSON(http.StatusOK, o.ResultFail("刷新失败"))
	}

	if result == nil || result.Code != "0" {
		return c.JSON(http.StatusOK, o.ResultFail("刷新失败"))
	}
	return c.JSON(http.StatusOK, o.ResultSuccess(result.Data))

}
