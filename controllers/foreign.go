package controllers

import (
	"github.com/astaxie/beego"
	"github.com/guotie/gogb2312"
	"io/ioutil"
	"strings"
	"time"
)

type ForeignController struct {
	beego.Controller
}

func (this *ForeignController) Get() {
	this.Data["Email"] = "z.geomantic@gmail.com"
	this.TplNames = "index.tpl"
	filePath, err := ioutil.ReadFile("cache")
	if err != nil {
		this.Data["Error"] = err.Error()
		return
	}

	content, err := ioutil.ReadFile(string(filePath))
	if err != nil {
		this.Data["Error"] = err.Error()
		return
	}
	output, _, _, _ := gogb2312.ConvertGB2312(content)
	code := strings.Split(string(output), "\n")
	this.Data["Code"] = code
	this.Data["Time"] = time.Now()
}
