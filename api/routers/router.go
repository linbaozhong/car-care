// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/linbaozhong/car-care/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/dict",
			beego.NSInclude(
				&controllers.Dict{},
			),
		),
		beego.NSNamespace("/brand",
			beego.NSInclude(
				&controllers.Brand{},
			),
		),
		beego.NSNamespace("/series",
			beego.NSInclude(
				&controllers.Series{},
			),
		),
		beego.NSNamespace("/vis",
			beego.NSInclude(
				&controllers.Visitor{},
			),
		),
		beego.NSNamespace("/clerk",
			beego.NSInclude(
				&controllers.Clerk{},
			),
		),
	)
	beego.AddNamespace(ns)
}
