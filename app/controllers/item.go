package controllers

import "github.com/revel/revel"

type Item struct {
	*revel.Controller
}

func (c Item) Index() revel.Result {


title := "Register"
	return c.Render(title)
	//#return c.Render()
}
