package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 拓展我们要写入token的信息
type Claims struct {
	Uid int `json:"uid"` // 扩展写入用户的id
	UserName string  `json:"username"`
	jwt.StandardClaims
}

func GenToken(uid int,username string) (string, error) {
	secretKey := "secretKey-HRforce2015" // 加密的key
	expiresIn := time.Duration(24 * 1) // 设置过期时间
	claims := Claims{
		Uid: uid,
		UserName:username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour.Truncate(expiresIn)).Unix(),
			Issuer:    "HRforce", // 签发的用户(自定义名字)
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用sha256进行加密claims这个结构体
	token, err := tokenClaims.SignedString([]byte(secretKey)) // 生成签名
	return token, err
}



