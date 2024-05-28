package global

import (
	"common_pkg/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CMP_VP     *viper.Viper
	CMP_LOG    *zap.Logger
	CMP_CONFIG config.Server
	CMP_DB     *gorm.DB
	CMP_DBMAP  map[string]*gorm.DB
)
