package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"test/LoveHome/models"
)

type UserControllers struct {
	beego.Controller
}

func (this *UserControllers) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (c *UserControllers) Reg() {

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer c.RetData(resp)

	var regRequestMap = make(map[string]interface{})

	json.Unmarshal(c.Ctx.Input.RequestBody, &regRequestMap)

	beego.Info("mobile = ", regRequestMap["mobile"])
	beego.Info("password = ", regRequestMap["password"])
	beego.Info("sms_code = ", regRequestMap["sms_code"])

	if regRequestMap["mobile"] == "" || regRequestMap["password"] == "" || regRequestMap["sms_code"] == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["ermsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	user := models.User{}
	user.Mobile = regRequestMap["mobile"].(string)
	user.Password_hash = regRequestMap["password"].(string)
	user.Name = regRequestMap["mobile"].(string)

	o := orm.NewOrm()

	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("insert error = ", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DATAERR)
		return
	}
	c.SetSession("name", user.Name)
	c.SetSession("user_id", id)
	c.SetSession("mobile", user.Mobile)

	return
}

func (c *UserControllers) Login() {

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer c.RetData(resp)

	var LoginRequestMap = make(map[string]interface{})

	json.Unmarshal(c.Ctx.Input.RequestBody, &LoginRequestMap)

	beego.Info("mobile = ", LoginRequestMap["mobile"])
	beego.Info("password = ", LoginRequestMap["password"])

	if LoginRequestMap["mobile"] == "" || LoginRequestMap["password"] == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["ermsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	var user models.User

	o := orm.NewOrm()

	qs := o.QueryTable("user")

	err := qs.Filter("mobile", LoginRequestMap["mobile"]).One(&user)
	if err != nil {
		beego.Info("insert error = ", err)
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}

	if user.Password_hash != LoginRequestMap["password"].(string) {
		resp["errno"] = models.RECODE_PWDERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		return
	}

	beego.Info("===================== login succ !!!! =============")

	c.SetSession("name", user.Name)
	c.SetSession("user_id", user.Id)
	c.SetSession("mobile", user.Mobile)
	return
}
