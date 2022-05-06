package configs

import (
	"fmt"

	"github.com/KwokBy/easy-ops/pkg/file"
	"github.com/spf13/viper"
)

type Configs struct {
	DB  DB  `yaml:"db"`
	JWT JWT `yaml:"jwt"`
	Docker Docker `yaml:"docker"`
}

func New() *Configs {
	c := &Configs{}
	v := viper.New()
	filePath, err := file.GetFilePath("./configs.yaml", 10)
	if err != nil {
		return nil
	}
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		return nil
	}
	if err := v.Unmarshal(&c); err != nil {
		return nil
	}
	fmt.Printf("%v\n", c)
	return c
}
