package controllers

import (
	"bytes"
	"code.google.com/p/go.text/encoding/simplifiedchinese"
	"code.google.com/p/go.text/transform"
	"github.com/astaxie/beego"
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

	bytes_code, err := ioutil.ReadFile(path)
	if err != nil {
		this.Data["Error"] = err.Error()
		return
	}
	GBKbytesReader := bytes.NewReader(bytes_code)
	UTFbytesReader := transform.NewReader(GBKbytesReader, simplifiedchinese.GBK.NewDecoder())
	RawCode, err := ioutil.ReadAll(UTFbytesReader)
	if err != nil {
		this.Data["Error"] = err.Error()
		return
	}
	// this.Data["Error"] = "this is error"
	Code := strings.Split(string(RawCode), "\n")
	this.Data["Code"] = Code
	this.Data["Time"] = time.Now()
	if path != "favicon.ico" {
		ioutil.WriteFile("cache", []byte(path), os.ModeAppend)
	}

}
