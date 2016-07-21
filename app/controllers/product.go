package controllers

import "github.com/revel/revel"

import "fmt"


type Product struct {
	*revel.Controller
}



type Stuff struct {
	Foo string ` json:"foo" xml:"foo" `
	Bar int ` json:"bar" xml:"bar" `
}


func (c Product) Insert() revel.Result {

	fmt.Printf("", "toto")

	data := make(map[string]interface{})
	data["error"] = nil
	stuff := Stuff{Foo: "xyz", Bar: 999}
	data["stuff"] = stuff


	return c.RenderText("hello")


	// return c.RenderJson(data) // for the json response



}
