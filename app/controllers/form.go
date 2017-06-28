package controllers

import (
"github.com/revel/revel"
"letsgo/app/models"
"log"
)

type Forum struct {
	GorpController
}

func (c Forum) Index() revel.Result {
	return c.Render()
}

func (c Forum) Edit(id int) revel.Result {
	return c.Render()
}

func (c Forum) Detail(id int) revel.Result {

	article := models.Article{}
	err :=	Dbm.SelectOne(&article, "select * from Article where Id = ?", id)
	if err != nil{
		log.Fatal("eerror")
	}
	return c.Render(article)
}

func (c Forum) Add() revel.Result {
	return c.Render()
}

func (c Forum) Review() revel.Result {
	return c.Render()
}

func (c Forum) List() revel.Result {
	return c.Render()
}

