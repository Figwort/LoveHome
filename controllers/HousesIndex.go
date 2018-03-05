package controllers

import (
	"github.com/astaxie/beego"
	"test/LoveHome/models"
)

type HouseIndexController struct {
	beego.Controller
}

func (this *HouseIndexController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *HouseIndexController) GetHousesIndex() {
	beego.Info("============= /api/v1.0/houses/index succ ========")

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["eermsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)
}
