package k8s

import (
	"context"
	"fmt"
	"github.com/0xweb-3/gokube/global"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type PodApi struct {
}

func (e *PodApi) ExampleTest(c *gin.Context) {
	ctx := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, item := range list.Items {
		fmt.Println(item.Namespace)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
