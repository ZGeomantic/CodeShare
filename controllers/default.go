package controllers

import (
	"github.com/astaxie/beego"
	"github.com/guotie/gogb2312"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	// this.Data["Website"] = "beego.me"
	this.Data["Email"] = "z.geomantic@gmail.com"
	this.TplNames = "index.tpl"
	path := this.Ctx.Input.Param(":splat")

	bytes_code2, err := ioutil.ReadFile(path)
	if err != nil {
		this.Data["Error"] = err.Error()
	}
	// this.Data["Error"] = "this is error"

	output, _, _, _ := gogb2312.ConvertGB2312(bytes_code2)
	code := strings.Split(string(output), "\n")
	this.Data["Code"] = code
	this.Data["Time"] = time.Now()
	if path != "favicon.ico" {
		ioutil.WriteFile("cache", []byte(path), os.ModeAppend)
	}

}
