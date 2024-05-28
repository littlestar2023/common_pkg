package global

import (
	"common_pkg/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	CMP_VP     *viper.Viper
	CMP_LOG    *zap.Logger
	CMP_CONFIG config.Server
	CMP_DB     *gorm.DB
	CMP_DBMAP  map[string]*gorm.DB
	GVA_REDIS  *redis.Client
	lock       sync.RWMutex
)

func GetGlobalDBByDBName(dbname string) *gorm.DB {

	lock.RLock()
	defer lock.RUnlock()
	return CMP_DBMAP[dbname]
}
