package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"time"
	"math"

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
	datas := c.Service.GetAll()
	// set the model and render the view template.
	count := len(datas)

	countPage := 2
	// int64(3/2) 1
	// math.Ceil(1.2) 2
	// math.Floor(1.2) 1
	pageCount := math.Ceil(float64(count) / float64(countPage))
	pageIndex := 1

	// <a href="/admin?id=1">first</a>

	pageIndex, _ = c.Ctx.URLParamInt("id")

	countStart := countPage * (pageIndex - 1)
	datalist := c.Service.GetLimit(countPage, countStart)

	firstPage := false
	if pageIndex == 1 {
		firstPage = true
	}

	lastPage := false
	if pageIndex == int(pageCount) {
		lastPage = true
	}

	return mvc.View{
		Name: "admin/index.html", // view template
		Data: iris.Map{
			"Title":    "管理后台",
			"Datalist": datalist,
			"count": count,
			"pageCount": pageCount,
			"current": pageIndex,
			"firstPage": firstPage,
			"lastPage": lastPage,
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

// GetAdd /admin/add get
func (c *AdminController) GetAdd() mvc.Result {

	return mvc.View{
		Name: "admin/add.html", // 视图html
		Data: iris.Map{
			"Title": "管理后台",
		},
		Layout: "admin/layout.html", // 重写模板(不要跟前端的layout混用)
	}
}

// PostAdd /admin/add post
func (c *AdminController) PostAdd() mvc.Result {
	info := models.StarInfo{}

	err := c.Ctx.ReadForm(&info)
	if err != nil {
		log.Fatal(err)
	}

	info.SysUpdated = int(time.Now().Unix())

	c.Service.Create(&info)

	return mvc.Response{
		Path: "/admin/",
	}
}