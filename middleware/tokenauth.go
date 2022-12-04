package middleware

import (
	j "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"hrforceAdmin/module/jwt"
	o "hrforceAdmin/utils/apioutput"
	"net/http"
)


func TokenAuth(next echo.HandlerFunc) echo.HandlerFunc{

	return func(c echo.Context) error {

		cookie,err := c.Cookie("_hrforceAdmin")
		if err != nil{
			return c.JSON(http.StatusOK,o.TokenInvalid("请传入有效token"))
		}
		tokenStr := cookie.Value

		//if tokenStr == ""{
		//	return c.JSON(http.StatusOK,o.ResultFail("请传入有效token"))
		//}
		claims :=jwt.Claims{}
		token, err := j.ParseWithClaims(tokenStr, &claims, func(token *j.Token) (interface{}, error) {
			return []byte("secretKey-HRforce2015"), nil
		})
		if err != nil {
			return c.JSON(http.StatusOK,o.TokenInvalid("token invalid"))
		}
		if _,ok := token.Claims.(*jwt.Claims); ok && token.Valid {
			return next(c)
		}else{
			return c.JSON(http.StatusOK,o.TokenInvalid("token invalid"))
		}

	}
}