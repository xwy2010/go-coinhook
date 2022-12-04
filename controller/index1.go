package controller

import (
	"fmt"
	"net/http"

	o "github.com/xwy2010/go-coinhook/common/apioutput"

	"github.com/xwy2010/go-coinhook/model"

	"github.com/labstack/echo/v4"
)

// IndexMobile 测试用
func IndexMobile(c echo.Context) (err error) {
	mobile := c.FormValue("mobile")
	user := &model.Admin{}

	fmt.Print("user mobile:", user.Mobile)
	user, err = user.GetByMobile(mobile)
	fmt.Println("get data for mobile err", err)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, o.ResultSuccess(user))

}
