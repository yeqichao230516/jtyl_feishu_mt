package core

import "github.com/spf13/viper"

var v *viper.Viper

func NewViper() {
	v = viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		Logger.Fatalf("读取配置文件失败: %v", err)
	}
}

func GetString(key string) string {
	return v.GetString(key)
}

func GetInt(key string) int {
	return v.GetInt(key)
}
