package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/xuzhenglun/WikiHub/models"
	"log"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "upload.html"
}

func (this *MainController) Post() {
	this.TplName = "upload.html"

	var form models.Form
	form.Add("name", this.GetString("name"))
	form.Add("zhiwei", this.GetString("职位"))
	form.Add("jishu", this.GetString("技术"))
	form.Add("fangxiang", this.GetString("方向"))
	form.Add("fangshi", this.GetString("方式"))
	form.Add("ziyuan", this.GetString("资源"))
	form.Add("xinge", this.GetString("性格"))
	form.Add("aihao", this.GetString("爱好"))
	form.Add("fuli", this.GetString("福利"))
	form.Add("xuezhang", this.GetString("学长"))
	form.Add("detail", this.GetString("detail"))

	challenge := this.GetString("geetest_challenge")
	validate := this.GetString("geetest_validate")
	seccode := this.GetString("geetest_seccode")

	log.Println(challenge, validate, seccode)

	sdk := models.NewTest()

	ok, err := sdk.Validate(challenge, validate, seccode)
	if err != nil {
		log.Println(err)
	}

	if ok {
		fmt.Fprintf(this.Ctx.ResponseWriter, "Success!")
		Tc := models.InitClient()
		models.CreateNewFile(form, nil, Tc)
	} else {
		fmt.Fprintf(this.Ctx.ResponseWriter, "Failed!")
	}
}
