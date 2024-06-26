package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/littlestar2023/common_pkg/consts"
	"github.com/littlestar2023/common_pkg/global"
	"github.com/spf13/viper"
	"os"
)

func InitialViper() *viper.Viper {

	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()

	if config == "" {
		if configEnv := os.Getenv(consts.ConfigEnv); configEnv == "" {
			config = consts.ConfigDefaultFile
		}
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CMP_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CMP_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}

func InitialViperWithCustomize(customizeType interface{}) *viper.Viper {

	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()

	if config == "" {
		if configEnv := os.Getenv(consts.ConfigEnv); configEnv == "" {
			config = consts.ConfigDefaultFile
		}
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	global.CMP_CONFIG.CustomizeConfig = customizeType

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CMP_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CMP_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
