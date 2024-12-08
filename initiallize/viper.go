package initiallize

import (
	"github.com/0xweb-3/gokube/global"
	"github.com/spf13/viper"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile("config/config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&global.CONF)
	if err != nil {
		panic(err)
	}
}
