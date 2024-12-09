package global

import (
	"github.com/0xweb-3/gokube/config"
	"k8s.io/client-go/kubernetes"
)

var (
	CONF          config.Server
	KubeConfigSet *kubernetes.Clientset
)
