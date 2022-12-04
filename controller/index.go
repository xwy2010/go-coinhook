package controller

import (
	"net/http"

	o "github.com/xwy2010/go-coinhook/common/apioutput"

	"github.com/labstack/echo/v4"
)

// Index 接口首页
func Index(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, o.ResultSuccess(nil))

}
