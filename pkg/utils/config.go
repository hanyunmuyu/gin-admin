package utils

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"os"
	"path/filepath"
)

var v *viper.Viper

func init() {
	v = viper.New()
	Config()
	v.WatchConfig()
}
func Config() *viper.Viper {
	if dir, err := os.Getwd(); err == nil {
		v.SetConfigName("env")  //设置配置文件的名字
		v.AddConfigPath(dir)    //添加配置文件所在的路径
		v.SetConfigType("yaml") // or viper.SetConfigType("YAML")
		if err := v.ReadInConfig(); err == nil {
			return v
		} else {
			dir, _ := os.Executable()
			exPath := filepath.Dir(dir)
			v.AddConfigPath(exPath) //添加配置文件所在的路径
			if err := v.ReadInConfig(); err == nil {
				return v
			}
		}
	}

	return v
}
