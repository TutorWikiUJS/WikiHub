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
	c.TplNames = "upload.html"
}

func (this *MainController) Post() {
	this.TplNames = "upload.html"
	name := this.GetString("name")
	detail := this.GetString("detail")

	challenge := this.GetString("geetest_challenge")
	validate := this.GetString("geetest_validate")
	seccode := this.GetString("geetest_seccode")

	log.Println(challenge,validate,seccode)

	sdk := models.NewTest()

	ok, err := sdk.Validate(challenge, validate, seccode)
	if err != nil {
		log.Println(err)
	}

	if ok {
		fmt.Fprintf(this.Ctx.ResponseWriter, "Success!")
		Tc := models.InitClient()
		models.CreateNewFile(name, detail, nil, Tc)
	} else {
		fmt.Fprintf(this.Ctx.ResponseWriter, "Failed!")
	}

}
