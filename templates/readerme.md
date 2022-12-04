## 安装Xorm tools

<p>go get xorm.io/cmd/xorm</p>


## 生成model文件
    注意：在cmd下运行，不能在powershell下运行
    usage: xorm reverse [-s] driverName datasourceName tmplPath [generatedPath] [tableFilterReg]
<p>xorm reverse -s mysql stage:"stage19871206@RDS"@tcp(rm-bp1n13996784ftop2o.mysql.rds.aliyuncs.com:3306)/hrforce?charset=utf8mb4 ./templates/goxorm ./model admin</p>

    will generated go files in ./model directory

## 接口测试
<code>
http://192.168.2.109:1234/login?username=xwy2010&password=c5889a75478095e721aff0cc276ebdcf

http://127.0.0.1:1234/v1/B/getCompanyuser?companycode=TEST
http://127.0.0.1:1234/v1/B/getCompanyuser?Companycode=TEST&&CompanyName=DISC测试
http://192.168.2.109:1234/v1/B/getCompanyuser?CompanyName=DISC

http://192.168.2.109:1234/v1/B/UpdateCompanysUsers?CompanyCode=TEST1&&ID=102
http://127.0.0.1:1234/v1/B/DelCompanysUsers?id=3
http://127.0.0.1:1234/v1/B/CountCompanysUsers
http://127.0.0.1:1234/v1/B/getCompanyuser?companycode=hrforce&&ContactPhonenum=&&page=0
http://127.0.0.1:1234/v1/B/getCompanyuser?companyname=hrforce&&ContactPhonenum=&&page=1

670b14728ad9902aecba32e22fa4f6bd  #000000




</code>

## go build
<code>
set GOARCH=amd64
set GOOS=linux
</code>


    "[go]": {
    "editor.snippetSuggestions": "none",
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
        "source.organizeImports": true
    }

    Your workspace is misconfigured: 
    go [-e -json -compiled=true -test=true -export=false -deps=true -find=false -- ./]: exit status 1: go: cannot find main module; see 'go help modules'
. Please see https://github.com/golang/tools/blob/master/gopls/doc/troubleshooting.md for more information or file an issue (https://github.com/golang/go/issues/new) if you believe this is a mistake.