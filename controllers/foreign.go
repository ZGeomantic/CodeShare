package controllers

import (
	"bytes"
	"code.google.com/p/go.text/encoding/simplifiedchinese"
	"code.google.com/p/go.text/transform"
	"github.com/astaxie/beego"
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

	bytes_code, err := ioutil.ReadFile(string(filePath))
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
}
