package controllers

import (
	"github.com/revel/revel"
	"letsgo/app/models"
)

type Admin struct {
	GorpController
}

func (c Admin)Blog() revel.Result {
	return c.Render()
}

func (c Admin)BlogAdd(a models.Article) revel.Result {
	Dbm.Insert(a)
	return c.Render()
}