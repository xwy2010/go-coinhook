# GO-Admin-CoinHook
##系统环境
1、gin
2、Mysql
3、gorm

##使用说明
首先找到配置文件，config/settings.yml， 同时也可创建开发环境配置，只需将默认配置文件 config/settings.yml 复制到 config/settings.dev.yml 即可，或者直接使用默认配置文件，直接修改config/settings.yml即可。

###初始化
项目中支持使用命令方式初始化基本数据结构和基础数据。 可以方便的使用 migrate 命令进行项目数据库结构和数据初始化。如下操作：
```shell
# 初始化
# macOS or linux 下使用
$ ./go-admin migrate -c=config/settings.dev.yml

# ⚠️注意:windows 下使用
$ go-admin.exe migrate -c=config/settings.dev.yml
```