package api

import (
	"github.com/0xweb-3/gokube/api/example"
	"github.com/0xweb-3/gokube/api/k8s"
)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
	K8sPodApiGroup  k8s.PodApi
}

var ApiGroupApp = new(ApiGroup)
