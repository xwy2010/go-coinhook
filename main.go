package main

import (
	"github.com/xwy2010/go-admin-coinhook/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6

// @title CoinHook API
// @version 1.0.0
// @description 基于Gin + Gorm的前后端分离权限管理系统的接口文档
// @license.name MIT
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @contact.name   天天下宇
// @contact.email  wenyu.xv@gmail.com
func main() {
	cmd.Execute()
}
