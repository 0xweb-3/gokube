package api

import "github.com/0xweb-3/gokube/api/example"

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
