package router

import "github.com/0xweb-3/gokube/router/example"

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
}

var RouterGroupApp = new(RouterGroup)
