package filter

import (
	"authentication_backend/errorValues"
	"authentication_backend/models"

	"github.com/astaxie/beego/context"
)

var AuthenticationFilter = func(ctx *context.Context) {
	accessToken := ctx.Input.Header("Authorization")
	err := models.AuthenticationAccessToken(accessToken)
	if err != nil {
		if err == errorValues.NotRegisterUserError {
			ctx.ResponseWriter.WriteHeader(404)
			ctx.ResponseWriter.Write([]byte("404 NotRegisterUser\n"))
		} else if err == errorValues.AccessTokenExpireError {
			ctx.ResponseWriter.WriteHeader(401)
			ctx.ResponseWriter.Write([]byte("401 AccessTokenExpireError\n"))
		} else {
			ctx.ResponseWriter.WriteHeader(403)
			ctx.ResponseWriter.Write([]byte("403 Authentication Error\n"))
		}
	}

}
