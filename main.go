package main

import (
	_ "CodeShare/routers"
	"github.com/astaxie/beego"
	"github.com/donnie4w/dom4g"
	"io/ioutil"
)

func main() {
	config_byte, err := ioutil.ReadFile("D:/config.xml")
	if err == nil {
		config := string(config_byte)
		beego.Debug(string(config))
		config_xml, err := dom4g.LoadByXml(config)
		if err == nil {
			ip := config_xml.Node("service").Node("NormalFuncSvr").Node("Pipe").Node("Sink").Node("Address").Value
			beego.Debug("ip is", ip)
		} else {
			beego.Debug("wrong", err)
		}
	}

	beego.Run()
}
