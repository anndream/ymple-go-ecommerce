package controllers

import "github.com/revel/revel"

import "fmt"

import "time"

import "github.com/patrickmn/go-cache"

type Item struct {
	*revel.Controller
}

func (c Item) Index() revel.Result {

	fmt.Println("Phone:", "toto")

	fmt.Println(c.Name)


	// exemple of get set key

	cache1 := cache.New(5 * time.Minute, 30 * time.Second)

	// Set the value of the key "foo" to "bar", with the default expiration time
	cache1.Set("foo", "bar", cache.DefaultExpiration)


	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	cache1.Set("baz", 42, cache.NoExpiration)


	// Get the string associated with the key "foo" from the cache
	foo, found := cache1.Get("foo")


	//foo2, found := cache1.Get("baz")

	if found {
		fmt.Println(foo)

	}

	var message string

	message += "new value"
	message += foo.(string)
	message += "end of new value"

	fmt.Println(message + " toto ")

	fmt.Println(foo)

	title2 := "Register"

	return c.Render(title2)
	//#return c.Render()
}
