package controllers

import (
	"authentication_backend/controllers/requests"
	"authentication_backend/controllers/responses"
	"authentication_backend/errorValues"
	"authentication_backend/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about AuthenticationController
type AuthenticationController struct {
	beego.Controller
}

// @Title Login
// @Description Login
// @Param	body		body 	requests.User	true		"body for user content"
// @Success 200 {int} responses.NewUserResponse
// @Failure 403 body is empty
// @Failure 404 body is empty
// @Failure 500 body is empty
// @router /login [post]
func (a *AuthenticationController) Login() {
	var req requests.User
	json.Unmarshal(a.Ctx.Input.RequestBody, &req)
	user, err := models.Login(req.UserName, req.Password)
	if err != nil {
		if err == errorValues.NotRegisterUserError {
			a.Ctx.Output.SetStatus(404)
			a.ServeJSON()
		}
		if err == errorValues.MissmatchPasswordError {
			a.Ctx.Output.SetStatus(403)
			a.ServeJSON()
		} else {
			a.Ctx.Output.SetStatus(500)
			a.ServeJSON()
		}
	}
	res := responses.NewUserResponse{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
	a.Data["json"] = res
	a.ServeJSON()
}

// @Title Signup
// @Description Signup
// @Param	body		body 	requests.User	true		"body for user content"
// @Success 200 {int} responses.NewUserResponse
// @Failure 409 body is empty
// @Failure 500 body is empty
// @router /signup [post]
func (a *AuthenticationController) Signup() {
	var req requests.User
	json.Unmarshal(a.Ctx.Input.RequestBody, &req)
	user, err := models.SignUp(req.UserName, req.Password)
	if err != nil {
		if err == errorValues.AlreadyExistUserError {
			a.Ctx.Output.SetStatus(409)
			a.ServeJSON()
		} else {
			a.Ctx.Output.SetStatus(500)
			a.ServeJSON()
		}
	}
	res := responses.NewUserResponse{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
	a.Data["json"] = res
	a.ServeJSON()
}

// @Title Logout
// @Description Logout
// @Param Authorization header string true "access token"
// @Success 201
// @router /logout [post]
func (a *AuthenticationController) Logout() {
	accessToken := a.Ctx.Input.Header("Authorization")

	models.Logout(accessToken)
	a.Ctx.Output.SetStatus(201)
	a.ServeJSON()
}

// @Title RefreshToken
// @Description RefreshToken
// @Param	body		body 	requests.RefreshToken	true		"body for RefreshToken content"
// @Success 200 responses.NewUserResponse
// @Failure 401 body is empty
// @router /refresh_token [post]
func (a *AuthenticationController) RefreshToken() {
	var req requests.RefreshToken
	json.Unmarshal(a.Ctx.Input.RequestBody, &req)
	user, err := models.RefreshToken(req.RefreshToken)
	if err != nil {
		a.Ctx.Output.SetStatus(401)
		a.ServeJSON()
	}
	res := responses.NewUserResponse{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
	a.Data["json"] = res
	a.ServeJSON()
}
