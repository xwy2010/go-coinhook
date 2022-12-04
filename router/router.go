package router

import (
	"fmt"
	"net/http"

	"github.com/xwy2010/go-coinhook/config"
	"github.com/xwy2010/go-coinhook/controller"
	"github.com/xwy2010/go-coinhook/middleware"
	"github.com/xwy2010/go-coinhook/model"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// Run 服务运行
func Run() {
	// Server
	e := echo.New()

	// Secure, XSS/CSS  HSTS
	e.Use(mw.SecureWithConfig(mw.DefaultSecureConfig))
	e.Use(mw.MethodOverride())

	// CORS
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	//e.POST("/indexmo",controller.IndexMobile)

	// test
	e.GET("/hello", func(c echo.Context) error {

		return c.String(http.StatusOK, "hello HRforce!")
	})
	// 用户模板文件
	e.File("/static/usersTemp.xlsx", "static/res/usersTemp.xlsx")
	// 表单校验
	e.Validator = &model.CustomValidator{}
	// 接口文档
	e.GET("/docs/*", echoSwagger.WrapHandler)
	// 接口
	g := e.Group("/v1", middleware.TokenAuth)
	g.POST("/index", controller.Index) // 接口首页测试

	// 启动服务
	srv := fmt.Sprintf("%s:%s", config.Config().SRV.IP, config.Config().SRV.Port)
	e.Logger.Fatal(e.Start(srv))

}
