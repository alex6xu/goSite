package controllers

import (
	"github.com/revel/revel"
	// "letsgo/app/models"
	"letsgo/app/models"
	"log"
)

type Blog struct {
	GorpController
}

func (c Blog) Index() revel.Result {
	return c.Render()
}

func (c Blog) Edit(id int) revel.Result {
	return c.Render()
}

func (c Blog) Detail(id int) revel.Result {

	article := models.Article{}
	err :=	Dbm.SelectOne(&article, "select * from Article where Id = ?", id)
	if err != nil{
		log.Fatal("eerror")
	}
	return c.Render(article)
}

func (c Blog) Add() revel.Result {
	return c.Render()
}

func (c Blog) Review() revel.Result {
	return c.Render()
}

func (c Blog) List() revel.Result {
	return c.Render()
}

