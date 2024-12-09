package main

import (
	"fmt"
	"github.com/0xweb-3/gokube/global"
	"github.com/0xweb-3/gokube/initiallize"
)

func main() {
	initiallize.Viper()
	r := initiallize.Routers()

	//initiallize.K8S()
	panic(r.Run(fmt.Sprintf("%s:%s", global.CONF.System.Addr, global.CONF.System.Port)))
}
