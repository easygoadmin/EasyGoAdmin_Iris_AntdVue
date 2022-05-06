// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package router

import (
	"easygoadmin/app/controller"
	"easygoadmin/app/middleware"
	"github.com/kataras/iris/v12"
	"github.com/rs/cors"
)

// 注册路由
func RegisterRouter(app *iris.Application) {
	// 登录验证中间件
	app.Use(middleware.CheckLogin)
	// 跨域解决方案
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		Debug:            true,
	})
	app.WrapRouter(crs.ServeHTTP)

	// 登录
	login := app.Party("/")
	{
		login.Get("/", controller.Login.Login)
		login.Post("/login", controller.Login.Login)
		login.Get("/captcha", controller.Login.Captcha)
		login.Put("/updateUserInfo", controller.Index.UpdateUserInfo)
		login.Put("/updatePwd", controller.Index.UpdatePwd)
		login.Get("/logout", controller.Index.Logout)
	}

	// 主页
	index := app.Party("/index")
	{
		index.Get("/menu", controller.Index.Menu)
		index.Get("/user", controller.Index.User)
	}

	// 文件上传
	upload := app.Party("/upload")
	{
		upload.Post("/uploadImage", controller.Upload.UploadImage)
	}

	// 职级管理
	level := app.Party("/level")
	{
		level.Get("/list", controller.Level.List)
		level.Get("/detail/{id:int}", controller.Level.Detail)
		level.Post("/add", controller.Level.Add)
		level.Put("/update", controller.Level.Update)
		level.Delete("/delete/{id:int}", controller.Level.Delete)
		level.Put("/status", controller.Level.Status)
		level.Get("/getLevelList", controller.Level.GetLevelList)
	}

	// 岗位管理
	position := app.Party("/position")
	{
		position.Get("/list", controller.Position.List)
		position.Get("/detail/{id:int}", controller.Position.Detail)
		position.Post("/add", controller.Position.Add)
		position.Put("/update", controller.Position.Update)
		position.Delete("/delete/{id:int}", controller.Position.Delete)
		position.Put("/status", controller.Position.Status)
		position.Get("/getPositionList", controller.Position.GetPositionList)
	}

	/* 角色路由 */
	role := app.Party("role")
	{
		role.Get("/list", controller.Role.List)
		role.Get("/detail/{id:int}", controller.Role.Detail)
		role.Post("/add", controller.Role.Add)
		role.Put("/update", controller.Role.Update)
		role.Delete("/delete/{id:int}", controller.Role.Delete)
		role.Put("/status", controller.Role.Status)
		role.Get("/getRoleList", controller.Role.GetRoleList)
	}

	/* 角色菜单权限 */
	roleMenu := app.Party("rolemenu")
	{
		roleMenu.Get("/index/{roleId:int}", controller.RoleMenu.Index)
		roleMenu.Post("/save", controller.RoleMenu.Save)
	}

	/* 部门管理 */
	dept := app.Party("dept")
	{
		dept.Get("/list", controller.Dept.List)
		dept.Get("/detail/{id:int}", controller.Dept.Detail)
		dept.Post("/add", controller.Dept.Add)
		dept.Put("/update", controller.Dept.Update)
		dept.Delete("/delete/{id:int}", controller.Dept.Delete)
		dept.Get("/getDeptList", controller.Dept.GetDeptList)
	}

	/* 用户管理 */
	user := app.Party("user")
	{
		user.Get("/list", controller.User.List)
		user.Get("/detail/{id:int}", controller.User.Detail)
		user.Post("/add", controller.User.Add)
		user.Put("/update", controller.User.Update)
		user.Delete("/delete/{id:int}", controller.User.Delete)
		user.Put("/status", controller.User.Status)
		user.Put("/resetPwd", controller.User.ResetPwd)
		user.Get("/checkUser/{username:string}", controller.User.CheckUser)
	}

	/* 菜单管理 */
	menu := app.Party("menu")
	{
		menu.Get("/list", controller.Menu.List)
		menu.Get("/detail/{id:int}", controller.Menu.Detail)
		menu.Post("/add", controller.Menu.Add)
		menu.Put("/update", controller.Menu.Update)
		menu.Delete("/delete/{id:int}", controller.Menu.Delete)
	}

	/* 友链管理 */
	link := app.Party("link")
	{
		link.Get("/list", controller.Link.List)
		link.Get("/detail/{id:int}", controller.Link.Detail)
		link.Post("/add", controller.Link.Add)
		link.Put("/update", controller.Link.Update)
		link.Delete("/delete/{id:int}", controller.Link.Delete)
		link.Put("/status", controller.Link.Status)
	}

	/* 城市管理 */
	city := app.Party("city")
	{
		city.Get("/list", controller.City.List)
		city.Get("/detail/{id:int}", controller.City.Detail)
		city.Post("/add", controller.City.Add)
		city.Put("/update", controller.City.Update)
		city.Delete("/delete/{id:int}", controller.City.Delete)
		city.Post("/getChilds", controller.City.GetChilds)
	}

	/* 通知管理 */
	notice := app.Party("notice")
	{
		notice.Get("/list", controller.Notice.List)
		notice.Get("/detail/{id:int}", controller.Notice.Detail)
		notice.Post("/add", controller.Notice.Add)
		notice.Put("/update", controller.Notice.Update)
		notice.Delete("/delete/{id:int}", controller.Notice.Delete)
		notice.Put("/status", controller.Notice.Status)
	}

	/* 会员等级 */
	memberlevel := app.Party("memberlevel")
	{
		memberlevel.Get("/list", controller.MemberLevel.List)
		memberlevel.Get("/detail/{id:int}", controller.MemberLevel.Detail)
		memberlevel.Post("/add", controller.MemberLevel.Add)
		memberlevel.Put("/update", controller.MemberLevel.Update)
		memberlevel.Delete("/delete/{id:int}", controller.MemberLevel.Delete)
		memberlevel.Get("/getMemberLevelList", controller.MemberLevel.GetMemberLevelList)
	}

	/* 会员管理 */
	member := app.Party("member")
	{
		member.Get("/list", controller.Member.List)
		member.Get("/detail/{id:int}", controller.Member.Detail)
		member.Post("/add", controller.Member.Add)
		member.Put("/update", controller.Member.Update)
		member.Delete("/delete/{id:int}", controller.Member.Delete)
		member.Put("/status", controller.Member.Status)
	}

	/* 字典管理 */
	dict := app.Party("dict")
	{
		dict.Get("/list", controller.Dict.List)
		dict.Get("/detail/{id:int}", controller.Dict.Detail)
		dict.Post("/add", controller.Dict.Add)
		dict.Put("/update", controller.Dict.Update)
		dict.Delete("/delete/{id:int}", controller.Dict.Delete)
	}

	/* 字典项管理 */
	dictdata := app.Party("dictdata")
	{
		dictdata.Get("/list", controller.DictData.List)
		dictdata.Get("/detail/{id:int}", controller.DictData.Detail)
		dictdata.Post("/add", controller.DictData.Add)
		dictdata.Put("/update", controller.DictData.Update)
		dictdata.Delete("/delete/{id:int}", controller.DictData.Delete)
	}

	/* 配置管理 */
	config := app.Party("config")
	{
		config.Get("/list", controller.Config.List)
		config.Get("/detail/{id:int}", controller.Config.Detail)
		config.Post("/add", controller.Config.Add)
		config.Put("/update", controller.Config.Update)
		config.Delete("/delete/{id:int}", controller.Config.Delete)
	}

	/* 配置项管理 */
	configdata := app.Party("configdata")
	{
		configdata.Get("/list", controller.ConfigData.List)
		configdata.Get("/detail/{id:int}", controller.ConfigData.Detail)
		configdata.Post("/add", controller.ConfigData.Add)
		configdata.Put("/update", controller.ConfigData.Update)
		configdata.Delete("/delete/{id:int}", controller.ConfigData.Delete)
		configdata.Put("/status", controller.ConfigData.Status)
	}

	/* 网站设置 */
	configweb := app.Party("configweb")
	{
		configweb.Get("/index", controller.ConfigWeb.Index)
		configweb.Put("/save", controller.ConfigWeb.Save)
	}

	/* 站点管理 */
	item := app.Party("item")
	{
		item.Get("/list", controller.Item.List)
		item.Get("/detail/{id:int}", controller.Item.Detail)
		item.Post("/add", controller.Item.Add)
		item.Put("/update", controller.Item.Update)
		item.Delete("/delete/{id:int}", controller.Item.Delete)
		item.Put("/status", controller.Item.Status)
		item.Get("/getItemList", controller.Item.GetItemList)
	}

	/* 栏目管理 */
	itemcate := app.Party("itemcate")
	{
		itemcate.Get("/list", controller.ItemCate.List)
		itemcate.Get("/detail/{id:int}", controller.ItemCate.Detail)
		itemcate.Post("/add", controller.ItemCate.Add)
		itemcate.Put("/update", controller.ItemCate.Update)
		itemcate.Delete("/delete/{id:int}", controller.ItemCate.Delete)
		itemcate.Get("/getCateList", controller.ItemCate.GetCateList)
	}

	/* 广告位管理 */
	adsort := app.Party("adsort")
	{
		adsort.Get("/list", controller.AdSort.List)
		adsort.Get("/detail/{id:int}", controller.AdSort.Detail)
		adsort.Post("/add", controller.AdSort.Add)
		adsort.Put("/update", controller.AdSort.Update)
		adsort.Delete("/delete/{id:int}", controller.AdSort.Delete)
		adsort.Get("/getAdSortList", controller.AdSort.GetAdSortList)
	}

	/* 广告管理 */
	ad := app.Party("ad")
	{
		ad.Get("/list", controller.Ad.List)
		ad.Get("/detail/{id:int}", controller.Ad.Detail)
		ad.Post("/add", controller.Ad.Add)
		ad.Put("/update", controller.Ad.Update)
		ad.Delete("/delete/{id:int}", controller.Ad.Delete)
		ad.Put("/status", controller.Ad.Status)
	}

	/* 代码生成器 */
	generate := app.Party("generate")
	{
		generate.Get("/list", controller.Generate.List)
		generate.Post("/generate", controller.Generate.Generate)
		generate.Post("/batchGenerate", controller.Generate.BatchGenerate)
	}

	/* 演示一 */
	example := app.Party("example")
	{
		example.Get("/list", controller.Example.List)
		example.Get("/detail/{id:int}", controller.Example.Detail)
		example.Post("/add", controller.Example.Add)
		example.Put("/update", controller.Example.Update)
		example.Delete("/delete/{id:int}", controller.Example.Delete)
		example.Put("/status", controller.Example.Status)
		example.Put("/setIsVip", controller.Example.IsVip)
	}

	/* 演示二 */
	example2 := app.Party("example2")
	{
		example2.Get("/list", controller.Example2.List)
		example2.Get("/detail/{id:int}", controller.Example2.Detail)
		example2.Post("/add", controller.Example2.Add)
		example2.Put("/update", controller.Example2.Update)
		example2.Delete("/delete/{id:int}", controller.Example2.Delete)
		example2.Put("/status", controller.Example2.Status)
	}

}
