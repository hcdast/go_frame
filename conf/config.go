package conf

import (
	beego "github.com/beego/beego/v2/server/web"
	"os"
)

// GetConf 获取配置变量
func GetConf(key string) string {
	value, _ := beego.AppConfig.String(key)
	if value == "" {
		value = os.Getenv(key)
	}
	return value
}