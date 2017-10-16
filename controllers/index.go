package controllers

import "github.com/astaxie/beego"

type IndexCotroller struct
{
	beego.Controller
}

func (this *IndexCotroller)Get(){
	this.TplName="index.html"
}

