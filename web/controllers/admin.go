package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"time"

	"superstar/models"
	"superstar/services"
)

// AdminController 嵌入 services.SuperstarService
// 調用services.SuperstarService的方法, 獲取數據
type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

// Get / get
func (c *AdminController) Get() mvc.Result {
	datalist := c.Service.GetAll()
	// set the model and render the view template.
	return mvc.View{
		Name: "admin/index.html", // view template
		Data: iris.Map{
			"Title":    "管理后台",
			"Datalist": datalist,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}

// GetEdit /admin/edit get
func (c *AdminController) GetEdit() mvc.Result {
	// /admin/edit?id={{.Id}}
	id, err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err == nil {
		data = c.Service.Get(id)
	} else {
		data = &models.StarInfo{
			Id: 0,
		}
	}
	//fmt.Println(id, data)
	// set the model and render the view template.
	return mvc.View{
		Name: "admin/edit.html", // 视图html
		Data: iris.Map{
			"Title": "管理后台",
			"info":  data,
		},
		Layout: "admin/layout.html", // 重写模板(不要跟前端的layout混用)
	}
}

// PostSave /admin/save post
func (c *AdminController) PostSave() mvc.Result {
	info := models.StarInfo{}
	// 读取表单("submit")
	err := c.Ctx.ReadForm(&info)
	//fmt.Printf("%v\n", info)
	if err != nil {
		log.Fatal(err)
	}
	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix())
		c.Service.Update(&info, []string{"name_zh", "name_en", "avatar",
			"birthday", "height", "weight", "club", "jersy", "coutry",
			"birthaddress", "feature", "moreinfo", "sys_updated"})
	} else {
		info.SysCreated = int(time.Now().Unix())
		c.Service.Create(&info)
	}
	// 跳转至 /admin
	return mvc.Response{
		Path: "/admin/",
	}
}

// GetDelete /admin/delete get
func (c *AdminController) GetDelete() mvc.Result {
	// /admin/delete?id={{.Id}}
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}
	// 跳转至 /admin
	return mvc.Response{
		Path: "/admin/",
	}
}
