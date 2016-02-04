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
	form.Add("职位", this.GetString("职位"))
	form.Add("实验室技术栈", this.GetString("技术"))
	form.Add("研究方向", this.GetString("方向"))
	form.Add("培养方式", this.GetString("方式"))
	form.Add("掌握的资源", this.GetString("资源"))
	form.Add("性格", this.GetString("性格"))
	form.Add("爱好", this.GetString("爱好"))
	form.Add("福利", this.GetString("福利"))
	form.Add("厉害的学长", this.GetString("学长"))
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
