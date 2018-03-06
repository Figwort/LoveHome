package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"test/LoveHome/models"
	"time"
)

type AreaControllers struct {
	beego.Controller
}

func (this *AreaControllers) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *AreaControllers) GetAreaInfo() {

	var areas []models.Area

	resp := make(map[string]interface{})

	resp["errno"] = 0
	resp["errmsg"] = "OK"

	defer c.RetData(resp)

	cache_conn, err := cache.NewCache("redis", `{"key":"lovehome","conn":"127.0.0.1:6379","dbNum":"0"}`)

	if err != nil {
		beego.Info("cache redis conn err, err = ", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	areas_info_val := cache_conn.Get("area_info")

	if areas_info_val != nil {
		beego.Info("================== get area_info from cache !!!=========")
		var area_info interface{}

		json.Unmarshal(areas_info_val.([]byte), &area_info)

		resp["data"] = area_info
		return
	}
	o := orm.NewOrm()

	qs := o.QueryTable("area")
	num, err := qs.All(&areas)

	if err != nil {
		resp["errno"] = models.RECODE_DBERR

		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	if num == 0 {
		resp["errno"] = models.RECODE_NODATA

		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	resp["data"] = areas

	areas_info_str, _ := json.Marshal(areas)

	if err := cache_conn.Put("area_info", areas_info_str, time.Second*3600); err != nil {
		beego.Info("set area_info -----> redis fail, err = ", err)

		resp["errno"] = models.RECODE_DBERR

		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

		return
	}

	return
}
