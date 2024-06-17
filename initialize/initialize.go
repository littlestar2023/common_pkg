package initialize

import (
	"github.com/littlestar2023/common_pkg/global"
)

func Initialize() {

	global.CMP_VP = InitialViper()
	global.CMP_LOG = InitialZap()
	if global.CMP_CONFIG.System.UseDb {
		global.CMP_DB = InitialGORM()
	}

	if global.CMP_CONFIG.System.UseRedis {
		global.GVA_REDIS = InitialRedis()
	}

	if global.CMP_CONFIG.System.UseNSQ {
		global.CMP_DBMAP = InitialDBMap()
	}
	return
}

func InitializeWithCustomize(customizeType interface{}) {

	global.CMP_VP = InitialViperWithCustomize(customizeType)
	global.CMP_LOG = InitialZap()
	if global.CMP_CONFIG.System.UseDb {
		global.CMP_DB = InitialGORM()
	}

	if global.CMP_CONFIG.System.UseRedis {
		global.GVA_REDIS = InitialRedis()
	}

	if global.CMP_CONFIG.System.UseNSQ {
		global.CMP_DBMAP = InitialDBMap()
	}
	return
}
