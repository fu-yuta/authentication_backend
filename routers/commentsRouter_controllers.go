package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["authentication_backend/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["authentication_backend/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["authentication_backend/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["authentication_backend/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["authentication_backend/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["authentication_backend/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "RefreshToken",
            Router: `/refresh_token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["authentication_backend/controllers:AuthenticationController"] = append(beego.GlobalControllerRouter["authentication_backend/controllers:AuthenticationController"],
        beego.ControllerComments{
            Method: "Signup",
            Router: `/signup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["authentication_backend/controllers:HelloController"] = append(beego.GlobalControllerRouter["authentication_backend/controllers:HelloController"],
        beego.ControllerComments{
            Method: "Hello",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
