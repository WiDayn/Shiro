package middleware

import (
	"Shiro/internal/config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

// JWTMiddleware 验证token是否有效并且增加username到上下文，如果没有则username为""(not nil)
func JWTMiddleware(ctx iris.Context) {
	var tokenString string

	// 首先尝试从Authorization头部获取token
	authHeader := ctx.GetHeader("Authorization")
	if authHeader != "" {
		tokenString = authHeader[len("Bearer "):] // 提取令牌
	} else {
		// 如果头部没有token，尝试从cookie获取
		tokenString = ctx.GetCookie("token")
	}

	ctx.Values().Set("username", "")

	// 如果没有找到token，继续
	if tokenString == "" {
		ctx.Next()
		return
	}

	// 解析和验证token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Application.JWTConfig.SigningKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Values().Set("username", claims["username"])
	}
	ctx.Next()
}

func JWTViewRequiredCheck(ctx iris.Context) {
	if ctx.Values().Get("username") == "" {
		ctx.Redirect("/admin/login", iris.StatusUnauthorized)
		return
	}
	ctx.Next()
}
