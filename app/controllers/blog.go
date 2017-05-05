package controllers

import (
	"github.com/revel/revel"
)

type Blog struct {
	*revel.Controller
}

func (c Blog) Index() revel.Result {
	return c.Render()
}

func (c Blog) Edit() revel.Result {
	return c.Render()
}

func (c Blog) Detail() revel.Result {
	return c.Render()
}

func (c Blog) Add() revel.Result {
	return c.Render()
}
