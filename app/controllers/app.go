package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Edit(id int) revel.Result {
	return c.Render()
}

func (c App) Detail(id int) revel.Result {
	return c.Render()
}

func (c App) Add() revel.Result {
	return c.Render()
}
