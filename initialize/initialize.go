package initialize

import (
	"common_pkg/core"
	"common_pkg/global"
)

func Initialize() (err error) {

	global.CMP_VP = core.Viper()
	global.CMP_LOG = core.Zap()

	return
}
