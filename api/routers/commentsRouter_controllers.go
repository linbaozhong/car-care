package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Brand"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Brand"],
		beego.ControllerComments{
			Method: "Gets",
			Router: `/gets`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Clerk"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Clerk"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Clerk"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Clerk"],
		beego.ControllerComments{
			Method: "UpdateInfo",
			Router: `/updateinfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"],
		beego.ControllerComments{
			Method: "Neishi",
			Router: `/neishi`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"],
		beego.ControllerComments{
			Method: "Qita",
			Router: `/qita`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"],
		beego.ControllerComments{
			Method: "Shangxian",
			Router: `/shangxian`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"],
		beego.ControllerComments{
			Method: "Waiguan",
			Router: `/waiguan`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Dict"],
		beego.ControllerComments{
			Method: "Yibiao",
			Router: `/yibiao`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Visitor"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Visitor"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Visitor"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Visitor"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
