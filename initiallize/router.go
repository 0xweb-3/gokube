package initiallize

import (
	"github.com/0xweb-3/gokube/router"
	"github.com/gin-gonic/gin"
)

func Routers() {
	r := gin.Default()
	exampleGroup := router.RouterGroupApp.ExampleRouterGroup
	exampleGroup.InitExample(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
