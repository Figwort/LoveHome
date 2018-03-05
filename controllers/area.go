package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"test/LoveHome/models"
)

type AreaControllers struct {
	beego.Controller
}

func (this *AreaControllers) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaControllers) GetAreaInfo() {

	o := orm.NewOrm()

	var areas []models.Area

	resp := make(map[string]interface{})

	resp["errno"] = 0
	resp["errmsg"] = "OK"

	defer c.RetData(resp)

	qs := o.QueryTable("area")
	num, err := qs.All(&areas)

	if err != nil {
		resp["errno"] = 4001

		resp["errmsg"] = "查询数据库失败"
		return
	}

	if num == 0 {
		resp["errno"] = 4002

		resp["errmsg"] = "没有数据"
		return
	}
	resp["data"] = areas
	return
}
