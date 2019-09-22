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

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Brand"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:Brand"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/linbaozhong/car-care/api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
